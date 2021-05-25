package gRpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SetUpServer(port string) *grpc.Server {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	log.Println("Running on port:", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return server
}

func AttachProtobufMetadata(statusCode int, field string, message string, description string) error {

	st := status.New(codes.Code(statusCode), message)
	fieldViolation := &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: description,
	}
	badRequest := &errdetails.BadRequest{}
	badRequest.FieldViolations = append(badRequest.FieldViolations, fieldViolation)

	st, err := st.WithDetails(badRequest)

	if err != nil {
		panic(fmt.Sprintf("Unexpected error attaching metadata: %v", err))
	}
	return st.Err()
}

func ExtractProtobufMetadata(err error) string {

	errorMessage := "An unexpected error occured."

	st := status.Convert(err)

	fmt.Println(st.WithDetails())
	for _, detail := range st.Details() {

		switch t := detail.(type) {
		case *errdetails.BadRequest:
			for _, violation := range t.GetFieldViolations() {
				errorMessage = violation.GetDescription()
			}
		}
	}

	return errorMessage
}
