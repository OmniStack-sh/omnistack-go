// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package omnistack_test

import (
	"context"
	"os"
	"testing"

	"github.com/OmniStack-sh/omnistack-go"
	"github.com/OmniStack-sh/omnistack-go/internal/testutil"
	"github.com/OmniStack-sh/omnistack-go/option"
	"github.com/OmniStack-sh/omnistack-go/shared"
)

func TestUsage(t *testing.T) {
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
	completion, err := client.Completions.New(context.TODO(), omnistack.CompletionNewParams{
		Model:  omnistack.F(omnistack.CompletionNewParamsModelGpt3_5TurboInstruct),
		Prompt: omnistack.F[omnistack.CompletionNewParamsPromptUnion](shared.UnionString("This is a test.")),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", completion.ID)
}
