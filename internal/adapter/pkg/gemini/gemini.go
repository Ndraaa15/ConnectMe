package gemini

import (
	"context"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/env"
	"github.com/google/generative-ai-go/genai"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

type Gemini struct {
	model *genai.GenerativeModel
}

type FoodScan struct {
	Text string `json:"text"`
}

func NewGemini(conf env.Gemini) *Gemini {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(conf.ApiKey))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create gemini client")
	}

	model := client.GenerativeModel(conf.Model)

	return &Gemini{
		model: model,
	}
}

func (g *Gemini) GenerateResponseForFoodScan(ctx context.Context, text string) (FoodScan, error) {
	_, err := g.model.GenerateContent(ctx, genai.Text("can you find the nutrition of this food!"), genai.ImageData("dasda", []byte("dsd")))
	if err != nil {
		return FoodScan{}, err
	}

	return FoodScan{}, nil
}
