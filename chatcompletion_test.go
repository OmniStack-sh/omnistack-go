// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package omnistack_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/omnistack-go"
	"github.com/stainless-sdks/omnistack-go/internal/testutil"
	"github.com/stainless-sdks/omnistack-go/option"
	"github.com/stainless-sdks/omnistack-go/shared"
)

func TestChatCompletionNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := omnistack.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Chats.Completions.New(context.TODO(), omnistack.ChatCompletionNewParams{
		Messages: omnistack.F([]omnistack.ChatCompletionNewParamsMessageUnion{omnistack.ChatCompletionNewParamsMessagesChatCompletionRequestSystemMessage{
			Content: omnistack.F[omnistack.ChatCompletionNewParamsMessagesChatCompletionRequestSystemMessageContentUnion](shared.UnionString("string")),
			Role:    omnistack.F(omnistack.ChatCompletionNewParamsMessagesChatCompletionRequestSystemMessageRoleSystem),
			Name:    omnistack.F("name"),
		}}),
		Model:            omnistack.F(omnistack.ChatCompletionNewParamsModelO1Preview),
		FrequencyPenalty: omnistack.F(-2.000000),
		FunctionCall:     omnistack.F[omnistack.ChatCompletionNewParamsFunctionCallUnion](omnistack.ChatCompletionNewParamsFunctionCallString(omnistack.ChatCompletionNewParamsFunctionCallStringNone)),
		Functions: omnistack.F([]omnistack.ChatCompletionNewParamsFunction{{
			Name:        omnistack.F("name"),
			Description: omnistack.F("description"),
			Parameters: omnistack.F(map[string]interface{}{
				"foo": "bar",
			}),
		}}),
		LogitBias: omnistack.F(map[string]int64{
			"foo": int64(0),
		}),
		Logprobs:            omnistack.F(true),
		MaxCompletionTokens: omnistack.F(int64(0)),
		MaxTokens:           omnistack.F(int64(0)),
		N:                   omnistack.F(int64(1)),
		ParallelToolCalls:   omnistack.F(true),
		PresencePenalty:     omnistack.F(-2.000000),
		ResponseFormat: omnistack.F[omnistack.ChatCompletionNewParamsResponseFormatUnion](omnistack.ChatCompletionNewParamsResponseFormatResponseFormatText{
			Type: omnistack.F(omnistack.ChatCompletionNewParamsResponseFormatResponseFormatTextTypeText),
		}),
		Seed:        omnistack.F(int64(-9007199254740991)),
		ServiceTier: omnistack.F(omnistack.ChatCompletionNewParamsServiceTierAuto),
		Stop:        omnistack.F[omnistack.ChatCompletionNewParamsStopUnion](shared.UnionString("string")),
		Stream:      omnistack.F(true),
		StreamOptions: omnistack.F(omnistack.ChatCompletionNewParamsStreamOptions{
			IncludeUsage: omnistack.F(true),
		}),
		Temperature: omnistack.F(1.000000),
		ToolChoice:  omnistack.F[omnistack.ChatCompletionNewParamsToolChoiceUnion](omnistack.ChatCompletionNewParamsToolChoiceString(omnistack.ChatCompletionNewParamsToolChoiceStringNone)),
		Tools: omnistack.F([]omnistack.ChatCompletionNewParamsTool{{
			Function: omnistack.F(omnistack.ChatCompletionNewParamsToolsFunction{
				Name:        omnistack.F("name"),
				Description: omnistack.F("description"),
				Parameters: omnistack.F(map[string]interface{}{
					"foo": "bar",
				}),
				Strict: omnistack.F(true),
			}),
			Type: omnistack.F(omnistack.ChatCompletionNewParamsToolsTypeFunction),
		}, {
			Function: omnistack.F(omnistack.ChatCompletionNewParamsToolsFunction{
				Name:        omnistack.F("name"),
				Description: omnistack.F("description"),
				Parameters: omnistack.F(map[string]interface{}{
					"foo": "bar",
				}),
				Strict: omnistack.F(true),
			}),
			Type: omnistack.F(omnistack.ChatCompletionNewParamsToolsTypeFunction),
		}, {
			Function: omnistack.F(omnistack.ChatCompletionNewParamsToolsFunction{
				Name:        omnistack.F("name"),
				Description: omnistack.F("description"),
				Parameters: omnistack.F(map[string]interface{}{
					"foo": "bar",
				}),
				Strict: omnistack.F(true),
			}),
			Type: omnistack.F(omnistack.ChatCompletionNewParamsToolsTypeFunction),
		}}),
		TopLogprobs: omnistack.F(int64(0)),
		TopP:        omnistack.F(1.000000),
		User:        omnistack.F("user-1234"),
	})
	if err != nil {
		var apierr *omnistack.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
