// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package quicksightiface provides an interface to enable mocking the Amazon QuickSight service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package quicksightiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/quicksight"
)

// QuickSightAPI provides an interface to enable mocking the
// quicksight.QuickSight service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon QuickSight.
//    func myFunc(svc quicksightiface.QuickSightAPI) bool {
//        // Make svc.CreateGroup request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := quicksight.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockQuickSightClient struct {
//        quicksightiface.QuickSightAPI
//    }
//    func (m *mockQuickSightClient) CreateGroup(input *quicksight.CreateGroupInput) (*quicksight.CreateGroupOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockQuickSightClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type QuickSightAPI interface {
	CreateGroup(*quicksight.CreateGroupInput) (*quicksight.CreateGroupOutput, error)
	CreateGroupWithContext(aws.Context, *quicksight.CreateGroupInput, ...request.Option) (*quicksight.CreateGroupOutput, error)
	CreateGroupRequest(*quicksight.CreateGroupInput) (*request.Request, *quicksight.CreateGroupOutput)

	CreateGroupMembership(*quicksight.CreateGroupMembershipInput) (*quicksight.CreateGroupMembershipOutput, error)
	CreateGroupMembershipWithContext(aws.Context, *quicksight.CreateGroupMembershipInput, ...request.Option) (*quicksight.CreateGroupMembershipOutput, error)
	CreateGroupMembershipRequest(*quicksight.CreateGroupMembershipInput) (*request.Request, *quicksight.CreateGroupMembershipOutput)

	DeleteGroup(*quicksight.DeleteGroupInput) (*quicksight.DeleteGroupOutput, error)
	DeleteGroupWithContext(aws.Context, *quicksight.DeleteGroupInput, ...request.Option) (*quicksight.DeleteGroupOutput, error)
	DeleteGroupRequest(*quicksight.DeleteGroupInput) (*request.Request, *quicksight.DeleteGroupOutput)

	DeleteGroupMembership(*quicksight.DeleteGroupMembershipInput) (*quicksight.DeleteGroupMembershipOutput, error)
	DeleteGroupMembershipWithContext(aws.Context, *quicksight.DeleteGroupMembershipInput, ...request.Option) (*quicksight.DeleteGroupMembershipOutput, error)
	DeleteGroupMembershipRequest(*quicksight.DeleteGroupMembershipInput) (*request.Request, *quicksight.DeleteGroupMembershipOutput)

	DeleteUser(*quicksight.DeleteUserInput) (*quicksight.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *quicksight.DeleteUserInput, ...request.Option) (*quicksight.DeleteUserOutput, error)
	DeleteUserRequest(*quicksight.DeleteUserInput) (*request.Request, *quicksight.DeleteUserOutput)

	DeleteUserByPrincipalId(*quicksight.DeleteUserByPrincipalIdInput) (*quicksight.DeleteUserByPrincipalIdOutput, error)
	DeleteUserByPrincipalIdWithContext(aws.Context, *quicksight.DeleteUserByPrincipalIdInput, ...request.Option) (*quicksight.DeleteUserByPrincipalIdOutput, error)
	DeleteUserByPrincipalIdRequest(*quicksight.DeleteUserByPrincipalIdInput) (*request.Request, *quicksight.DeleteUserByPrincipalIdOutput)

	DescribeGroup(*quicksight.DescribeGroupInput) (*quicksight.DescribeGroupOutput, error)
	DescribeGroupWithContext(aws.Context, *quicksight.DescribeGroupInput, ...request.Option) (*quicksight.DescribeGroupOutput, error)
	DescribeGroupRequest(*quicksight.DescribeGroupInput) (*request.Request, *quicksight.DescribeGroupOutput)

	DescribeUser(*quicksight.DescribeUserInput) (*quicksight.DescribeUserOutput, error)
	DescribeUserWithContext(aws.Context, *quicksight.DescribeUserInput, ...request.Option) (*quicksight.DescribeUserOutput, error)
	DescribeUserRequest(*quicksight.DescribeUserInput) (*request.Request, *quicksight.DescribeUserOutput)

	GetDashboardEmbedUrl(*quicksight.GetDashboardEmbedUrlInput) (*quicksight.GetDashboardEmbedUrlOutput, error)
	GetDashboardEmbedUrlWithContext(aws.Context, *quicksight.GetDashboardEmbedUrlInput, ...request.Option) (*quicksight.GetDashboardEmbedUrlOutput, error)
	GetDashboardEmbedUrlRequest(*quicksight.GetDashboardEmbedUrlInput) (*request.Request, *quicksight.GetDashboardEmbedUrlOutput)

	ListGroupMemberships(*quicksight.ListGroupMembershipsInput) (*quicksight.ListGroupMembershipsOutput, error)
	ListGroupMembershipsWithContext(aws.Context, *quicksight.ListGroupMembershipsInput, ...request.Option) (*quicksight.ListGroupMembershipsOutput, error)
	ListGroupMembershipsRequest(*quicksight.ListGroupMembershipsInput) (*request.Request, *quicksight.ListGroupMembershipsOutput)

	ListGroups(*quicksight.ListGroupsInput) (*quicksight.ListGroupsOutput, error)
	ListGroupsWithContext(aws.Context, *quicksight.ListGroupsInput, ...request.Option) (*quicksight.ListGroupsOutput, error)
	ListGroupsRequest(*quicksight.ListGroupsInput) (*request.Request, *quicksight.ListGroupsOutput)

	ListUserGroups(*quicksight.ListUserGroupsInput) (*quicksight.ListUserGroupsOutput, error)
	ListUserGroupsWithContext(aws.Context, *quicksight.ListUserGroupsInput, ...request.Option) (*quicksight.ListUserGroupsOutput, error)
	ListUserGroupsRequest(*quicksight.ListUserGroupsInput) (*request.Request, *quicksight.ListUserGroupsOutput)

	ListUsers(*quicksight.ListUsersInput) (*quicksight.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *quicksight.ListUsersInput, ...request.Option) (*quicksight.ListUsersOutput, error)
	ListUsersRequest(*quicksight.ListUsersInput) (*request.Request, *quicksight.ListUsersOutput)

	RegisterUser(*quicksight.RegisterUserInput) (*quicksight.RegisterUserOutput, error)
	RegisterUserWithContext(aws.Context, *quicksight.RegisterUserInput, ...request.Option) (*quicksight.RegisterUserOutput, error)
	RegisterUserRequest(*quicksight.RegisterUserInput) (*request.Request, *quicksight.RegisterUserOutput)

	UpdateGroup(*quicksight.UpdateGroupInput) (*quicksight.UpdateGroupOutput, error)
	UpdateGroupWithContext(aws.Context, *quicksight.UpdateGroupInput, ...request.Option) (*quicksight.UpdateGroupOutput, error)
	UpdateGroupRequest(*quicksight.UpdateGroupInput) (*request.Request, *quicksight.UpdateGroupOutput)

	UpdateUser(*quicksight.UpdateUserInput) (*quicksight.UpdateUserOutput, error)
	UpdateUserWithContext(aws.Context, *quicksight.UpdateUserInput, ...request.Option) (*quicksight.UpdateUserOutput, error)
	UpdateUserRequest(*quicksight.UpdateUserInput) (*request.Request, *quicksight.UpdateUserOutput)
}

var _ QuickSightAPI = (*quicksight.QuickSight)(nil)
