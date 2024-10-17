// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package omnistack

import (
	"context"
	"net/http"

	"github.com/OmniStack-sh/omnistack-go/internal/apijson"
	"github.com/OmniStack-sh/omnistack-go/internal/param"
	"github.com/OmniStack-sh/omnistack-go/internal/requestconfig"
	"github.com/OmniStack-sh/omnistack-go/option"
)

// CompletionService contains methods and other services that help with interacting
// with the omnistack API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompletionService] method instead.
type CompletionService struct {
	Options []option.RequestOption
}

// NewCompletionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCompletionService(opts ...option.RequestOption) (r *CompletionService) {
	r = &CompletionService{}
	r.Options = opts
	return
}

// Creates a completion for the provided prompt and parameters.
func (r *CompletionService) New(ctx context.Context, body CompletionNewParams, opts ...option.RequestOption) (res *CompletionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Represents a completion response from the API. Note: both the streamed and
// non-streamed response objects share the same shape (unlike the chat endpoint).
type CompletionNewResponse struct {
	// A unique identifier for the completion.
	ID string `json:"id,required"`
	// The list of completion choices the model generated for the input prompt.
	Choices []CompletionNewResponseChoice `json:"choices,required"`
	// The Unix timestamp (in seconds) of when the completion was created.
	Created int64 `json:"created,required"`
	// The model used for completion.
	Model string `json:"model,required"`
	// The object type, which is always "text_completion"
	Object CompletionNewResponseObject `json:"object,required"`
	// This fingerprint represents the backend configuration that the model runs with.
	//
	// Can be used in conjunction with the `seed` request parameter to understand when
	// backend changes have been made that might impact determinism.
	SystemFingerprint string `json:"system_fingerprint"`
	// Usage statistics for the completion request.
	Usage CompletionNewResponseUsage `json:"usage"`
	JSON  completionNewResponseJSON  `json:"-"`
}

// completionNewResponseJSON contains the JSON metadata for the struct
// [CompletionNewResponse]
type completionNewResponseJSON struct {
	ID                apijson.Field
	Choices           apijson.Field
	Created           apijson.Field
	Model             apijson.Field
	Object            apijson.Field
	SystemFingerprint apijson.Field
	Usage             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CompletionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionNewResponseJSON) RawJSON() string {
	return r.raw
}

type CompletionNewResponseChoice struct {
	// The reason the model stopped generating tokens. This will be `stop` if the model
	// hit a natural stop point or a provided stop sequence, `length` if the maximum
	// number of tokens specified in the request was reached, or `content_filter` if
	// content was omitted due to a flag from our content filters.
	FinishReason CompletionNewResponseChoicesFinishReason `json:"finish_reason,required"`
	Index        int64                                    `json:"index,required"`
	Logprobs     CompletionNewResponseChoicesLogprobs     `json:"logprobs,required,nullable"`
	Text         string                                   `json:"text,required"`
	JSON         completionNewResponseChoiceJSON          `json:"-"`
}

// completionNewResponseChoiceJSON contains the JSON metadata for the struct
// [CompletionNewResponseChoice]
type completionNewResponseChoiceJSON struct {
	FinishReason apijson.Field
	Index        apijson.Field
	Logprobs     apijson.Field
	Text         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CompletionNewResponseChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionNewResponseChoiceJSON) RawJSON() string {
	return r.raw
}

// The reason the model stopped generating tokens. This will be `stop` if the model
// hit a natural stop point or a provided stop sequence, `length` if the maximum
// number of tokens specified in the request was reached, or `content_filter` if
// content was omitted due to a flag from our content filters.
type CompletionNewResponseChoicesFinishReason string

const (
	CompletionNewResponseChoicesFinishReasonStop          CompletionNewResponseChoicesFinishReason = "stop"
	CompletionNewResponseChoicesFinishReasonLength        CompletionNewResponseChoicesFinishReason = "length"
	CompletionNewResponseChoicesFinishReasonContentFilter CompletionNewResponseChoicesFinishReason = "content_filter"
)

func (r CompletionNewResponseChoicesFinishReason) IsKnown() bool {
	switch r {
	case CompletionNewResponseChoicesFinishReasonStop, CompletionNewResponseChoicesFinishReasonLength, CompletionNewResponseChoicesFinishReasonContentFilter:
		return true
	}
	return false
}

type CompletionNewResponseChoicesLogprobs struct {
	TextOffset    []int64                                  `json:"text_offset"`
	TokenLogprobs []float64                                `json:"token_logprobs"`
	Tokens        []string                                 `json:"tokens"`
	TopLogprobs   []map[string]float64                     `json:"top_logprobs"`
	JSON          completionNewResponseChoicesLogprobsJSON `json:"-"`
}

// completionNewResponseChoicesLogprobsJSON contains the JSON metadata for the
// struct [CompletionNewResponseChoicesLogprobs]
type completionNewResponseChoicesLogprobsJSON struct {
	TextOffset    apijson.Field
	TokenLogprobs apijson.Field
	Tokens        apijson.Field
	TopLogprobs   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CompletionNewResponseChoicesLogprobs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionNewResponseChoicesLogprobsJSON) RawJSON() string {
	return r.raw
}

// The object type, which is always "text_completion"
type CompletionNewResponseObject string

const (
	CompletionNewResponseObjectTextCompletion CompletionNewResponseObject = "text_completion"
)

func (r CompletionNewResponseObject) IsKnown() bool {
	switch r {
	case CompletionNewResponseObjectTextCompletion:
		return true
	}
	return false
}

// Usage statistics for the completion request.
type CompletionNewResponseUsage struct {
	// Number of tokens in the generated completion.
	CompletionTokens int64 `json:"completion_tokens,required"`
	// Number of tokens in the prompt.
	PromptTokens int64 `json:"prompt_tokens,required"`
	// Total number of tokens used in the request (prompt + completion).
	TotalTokens int64 `json:"total_tokens,required"`
	// Breakdown of tokens used in a completion.
	CompletionTokensDetails CompletionNewResponseUsageCompletionTokensDetails `json:"completion_tokens_details"`
	JSON                    completionNewResponseUsageJSON                    `json:"-"`
}

// completionNewResponseUsageJSON contains the JSON metadata for the struct
// [CompletionNewResponseUsage]
type completionNewResponseUsageJSON struct {
	CompletionTokens        apijson.Field
	PromptTokens            apijson.Field
	TotalTokens             apijson.Field
	CompletionTokensDetails apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CompletionNewResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionNewResponseUsageJSON) RawJSON() string {
	return r.raw
}

// Breakdown of tokens used in a completion.
type CompletionNewResponseUsageCompletionTokensDetails struct {
	// Tokens generated by the model for reasoning.
	ReasoningTokens int64                                                 `json:"reasoning_tokens"`
	JSON            completionNewResponseUsageCompletionTokensDetailsJSON `json:"-"`
}

// completionNewResponseUsageCompletionTokensDetailsJSON contains the JSON metadata
// for the struct [CompletionNewResponseUsageCompletionTokensDetails]
type completionNewResponseUsageCompletionTokensDetailsJSON struct {
	ReasoningTokens apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CompletionNewResponseUsageCompletionTokensDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionNewResponseUsageCompletionTokensDetailsJSON) RawJSON() string {
	return r.raw
}

type CompletionNewParams struct {
	// ID of the model to use. You can use the
	// [List models](/docs/api-reference/models/list) API to see all of your available
	// models, or see our [Model overview](/docs/models/overview) for descriptions of
	// them.
	Model param.Field[CompletionNewParamsModel] `json:"model,required"`
	// The prompt(s) to generate completions for, encoded as a string, array of
	// strings, array of tokens, or array of token arrays.
	//
	// Note that <|endoftext|> is the document separator that the model sees during
	// training, so if a prompt is not specified the model will generate as if from the
	// beginning of a new document.
	Prompt param.Field[CompletionNewParamsPromptUnion] `json:"prompt,required"`
	// Generates `best_of` completions server-side and returns the "best" (the one with
	// the highest log probability per token). Results cannot be streamed.
	//
	// When used with `n`, `best_of` controls the number of candidate completions and
	// `n` specifies how many to return – `best_of` must be greater than `n`.
	//
	// **Note:** Because this parameter generates many completions, it can quickly
	// consume your token quota. Use carefully and ensure that you have reasonable
	// settings for `max_tokens` and `stop`.
	BestOf param.Field[int64] `json:"best_of"`
	// Echo back the prompt in addition to the completion
	Echo param.Field[bool] `json:"echo"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their
	// existing frequency in the text so far, decreasing the model's likelihood to
	// repeat the same line verbatim.
	//
	// [See more information about frequency and presence penalties.](/docs/guides/text-generation/parameter-details)
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// Modify the likelihood of specified tokens appearing in the completion.
	//
	// Accepts a JSON object that maps tokens (specified by their token ID in the GPT
	// tokenizer) to an associated bias value from -100 to 100. You can use this
	// [tokenizer tool](/tokenizer?view=bpe) to convert text to token IDs.
	// Mathematically, the bias is added to the logits generated by the model prior to
	// sampling. The exact effect will vary per model, but values between -1 and 1
	// should decrease or increase likelihood of selection; values like -100 or 100
	// should result in a ban or exclusive selection of the relevant token.
	//
	// As an example, you can pass `{"50256": -100}` to prevent the <|endoftext|> token
	// from being generated.
	LogitBias param.Field[map[string]int64] `json:"logit_bias"`
	// Include the log probabilities on the `logprobs` most likely output tokens, as
	// well the chosen tokens. For example, if `logprobs` is 5, the API will return a
	// list of the 5 most likely tokens. The API will always return the `logprob` of
	// the sampled token, so there may be up to `logprobs+1` elements in the response.
	//
	// The maximum value for `logprobs` is 5.
	Logprobs param.Field[int64] `json:"logprobs"`
	// The maximum number of [tokens](/tokenizer) that can be generated in the
	// completion.
	//
	// The token count of your prompt plus `max_tokens` cannot exceed the model's
	// context length.
	// [Example Python code](https://cookbook.openai.com/examples/how_to_count_tokens_with_tiktoken)
	// for counting tokens.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// How many completions to generate for each prompt.
	//
	// **Note:** Because this parameter generates many completions, it can quickly
	// consume your token quota. Use carefully and ensure that you have reasonable
	// settings for `max_tokens` and `stop`.
	N param.Field[int64] `json:"n"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on
	// whether they appear in the text so far, increasing the model's likelihood to
	// talk about new topics.
	//
	// [See more information about frequency and presence penalties.](/docs/guides/text-generation/parameter-details)
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If specified, our system will make a best effort to sample deterministically,
	// such that repeated requests with the same `seed` and parameters should return
	// the same result.
	//
	// Determinism is not guaranteed, and you should refer to the `system_fingerprint`
	// response parameter to monitor changes in the backend.
	Seed param.Field[int64] `json:"seed"`
	// Up to 4 sequences where the API will stop generating further tokens. The
	// returned text will not contain the stop sequence.
	Stop param.Field[CompletionNewParamsStopUnion] `json:"stop"`
	// Whether to stream back partial progress. If set, tokens will be sent as
	// data-only
	// [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format)
	// as they become available, with the stream terminated by a `data: [DONE]`
	// message.
	// [Example Python code](https://cookbook.openai.com/examples/how_to_stream_completions).
	Stream param.Field[bool] `json:"stream"`
	// Options for streaming response. Only set this when you set `stream: true`.
	StreamOptions param.Field[CompletionNewParamsStreamOptions] `json:"stream_options"`
	// The suffix that comes after a completion of inserted text.
	//
	// This parameter is only supported for `gpt-3.5-turbo-instruct`.
	Suffix param.Field[string] `json:"suffix"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will
	// make the output more random, while lower values like 0.2 will make it more
	// focused and deterministic.
	//
	// We generally recommend altering this or `top_p` but not both.
	Temperature param.Field[float64] `json:"temperature"`
	// An alternative to sampling with temperature, called nucleus sampling, where the
	// model considers the results of the tokens with top_p probability mass. So 0.1
	// means only the tokens comprising the top 10% probability mass are considered.
	//
	// We generally recommend altering this or `temperature` but not both.
	TopP param.Field[float64] `json:"top_p"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor
	// and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).
	User param.Field[string] `json:"user"`
}

func (r CompletionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CompletionNewParamsModel string

const (
	CompletionNewParamsModelGpt3_5TurboInstruct CompletionNewParamsModel = "gpt-3.5-turbo-instruct"
	CompletionNewParamsModelDavinci002          CompletionNewParamsModel = "davinci-002"
	CompletionNewParamsModelBabbage002          CompletionNewParamsModel = "babbage-002"
)

func (r CompletionNewParamsModel) IsKnown() bool {
	switch r {
	case CompletionNewParamsModelGpt3_5TurboInstruct, CompletionNewParamsModelDavinci002, CompletionNewParamsModelBabbage002:
		return true
	}
	return false
}

// The prompt(s) to generate completions for, encoded as a string, array of
// strings, array of tokens, or array of token arrays.
//
// Note that <|endoftext|> is the document separator that the model sees during
// training, so if a prompt is not specified the model will generate as if from the
// beginning of a new document.
//
// Satisfied by [shared.UnionString], [CompletionNewParamsPromptArray],
// [CompletionNewParamsPromptArray], [CompletionNewParamsPromptArray].
type CompletionNewParamsPromptUnion interface {
	ImplementsCompletionNewParamsPromptUnion()
}

type CompletionNewParamsPromptArray []string

func (r CompletionNewParamsPromptArray) ImplementsCompletionNewParamsPromptUnion() {}

// Up to 4 sequences where the API will stop generating further tokens. The
// returned text will not contain the stop sequence.
//
// Satisfied by [shared.UnionString], [CompletionNewParamsStopArray].
type CompletionNewParamsStopUnion interface {
	ImplementsCompletionNewParamsStopUnion()
}

type CompletionNewParamsStopArray []string

func (r CompletionNewParamsStopArray) ImplementsCompletionNewParamsStopUnion() {}

// Options for streaming response. Only set this when you set `stream: true`.
type CompletionNewParamsStreamOptions struct {
	// If set, an additional chunk will be streamed before the `data: [DONE]` message.
	// The `usage` field on this chunk shows the token usage statistics for the entire
	// request, and the `choices` field will always be an empty array. All other chunks
	// will also include a `usage` field, but with a null value.
	IncludeUsage param.Field[bool] `json:"include_usage"`
}

func (r CompletionNewParamsStreamOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
