package push

//
//import (
//	"context"
//	"github.com/google/go-cmp/cmp/cmpopts"
//	"testing"
//
//	"github.com/aws/aws-sdk-go-v2/aws"
//	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
//	"github.com/google/go-cmp/cmp"
//)
//
//type mockPublisher struct {
//	body  string
//	attrs map[string]types.MessageAttributeValue
//}
//
//func (m *mockPublisher) PublishMailMessage(ctx context.Context, body string, attrs map[string]types.MessageAttributeValue) error {
//	m.body = body
//	m.attrs = attrs
//	return nil
//}
//func (m *mockPublisher) GetMailMessage(ctx context.Context) ([]types.Message, error) { return nil, nil }
//
//func TestPublishMessageUseCase_Do(t *testing.T) {
//	repo := &mockPublisher{}
//	uc := NewPublishMessageUseCase(repo)
//
//	err := uc.Do(context.Background(), "body", "Bob", "bob@example.com", "hello")
//	if err != nil {
//		t.Fatalf("unexpected error: %v", err)
//	}
//
//	if diff := cmp.Diff("{\"title\":\"aa\",\"message\":\"\\u003c!DOCTYPE html\\u003e\\u003chtml\\u003e\\u003cbody\\u003e\\u003ch2\\u003eaa\\u003c/h2\\u003e\\u003cimg src='https://bucket.s3.ap-northeast-1.amazonaws.com/keyname' /\\u003e\\u003c/body\\u003e\\u003c/html\\u003e\"}", repo.body); diff != "" {
//		t.Errorf("body mismatch (-want +got):\n%s", diff)
//	}
//
//	expected := map[string]types.MessageAttributeValue{
//		"userName": {DataType: aws.String("String"), StringValue: aws.String("Bob")},
//		"message":  {DataType: aws.String("String"), StringValue: aws.String("hello")},
//		"address":  {DataType: aws.String("String"), StringValue: aws.String("bob@example.com")},
//	}
//	if diff := cmp.Diff(expected, repo.attrs,
//		cmpopts.IgnoreUnexported(types.MessageAttributeValue{}),
//	); diff != "" {
//		t.Errorf("message attributes mismatch (-want +got):\n%s", diff)
//	}
//
//}
//
//func BenchmarkPublishMessageUseCase_Do(b *testing.B) {
//	repo := &mockPublisher{}
//	uc := NewPublishMessageUseCase(repo)
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		uc.Do(context.Background(), "hoge", "hoge", "hoge", "hoge")
//	}
//}
