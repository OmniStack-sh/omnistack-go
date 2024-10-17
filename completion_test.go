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

func TestCompletionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Completions.New(context.TODO(), omnistack.CompletionNewParams{
		Model:            omnistack.F(omnistack.CompletionNewParamsModelGpt3_5TurboInstruct),
		Prompt:           omnistack.F[omnistack.CompletionNewParamsPromptUnion](shared.UnionString("This is a test.")),
		BestOf:           omnistack.F(int64(0)),
		Echo:             omnistack.F(true),
		FrequencyPenalty: omnistack.F(-2.000000),
		LogitBias: omnistack.F(map[string]int64{
			"foo": int64(0),
		}),
		Logprobs:        omnistack.F(int64(0)),
		MaxTokens:       omnistack.F(int64(16)),
		N:               omnistack.F(int64(1)),
		PresencePenalty: omnistack.F(-2.000000),
		Seed:            omnistack.F(int64(-9007199254740991)),
		Stop:            omnistack.F[omnistack.CompletionNewParamsStopUnion](shared.UnionString("\n")),
		Stream:          omnistack.F(true),
		StreamOptions: omnistack.F(omnistack.CompletionNewParamsStreamOptions{
			IncludeUsage: omnistack.F(true),
		}),
		Suffix:      omnistack.F("test."),
		Temperature: omnistack.F(1.000000),
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
