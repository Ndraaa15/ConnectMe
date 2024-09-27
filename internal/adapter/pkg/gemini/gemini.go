package gemini

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/env"
	"github.com/Ndraaa15/ConnectMe/internal/core/dto"
	"github.com/google/generative-ai-go/genai"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

type Gemini struct {
	model *genai.GenerativeModel
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

func (g *Gemini) GenerateResponseForProblem(ctx context.Context, text string, picture []byte) (dto.ResponseProblem, error) {
	genaiPart := []genai.Part{
		genai.Text("Anda adalah sebuah bot AI yang dirancang untuk membantu menyelesaikan berbagai masalah dengan memberikan solusi yang tepat."),
		genai.Text("Berikut adalah masalah yang disampaikan oleh pengguna:"),
		genai.Text(text),
		genai.Text("Tugas Anda adalah memberikan solusi yang jelas dan ringkas terkait masalah ini. Solusi harus praktis dan dapat diterapkan."),
		genai.Text("Selain solusi, berikan juga kata kunci yang relevan dengan masalah dan solusi yang telah Anda berikan."),
		genai.Text("Berikan respons Anda dalam format JSON dengan struktur berikut:"),
		genai.Text(`{
        "solution": "<solusi yang Anda berikan>",
        "keyword": ["<kata kunci 1>", "<kata kunci 2>", "..."]
    }`),
		genai.Text("Pastikan bahwa setiap kata kunci adalah kata tunggal yang menggambarkan masalah dan solusi dengan baik."),
	}

	if picture != nil {
		genaiPart = append(genaiPart, genai.Text("Ini merupakan gambar tambahan untuk membantu anda memberikan solusi"))
		genaiPart = append(genaiPart, genai.ImageData("jpg", picture))
	}

	content, err := g.model.GenerateContent(ctx, genaiPart...)
	if err != nil {
		return dto.ResponseProblem{}, err
	}

	part := content.Candidates[0].Content.Parts[0]
	jsonByte, err := json.Marshal(part)
	if err != nil {
		return dto.ResponseProblem{}, nil
	}

	jsonStr, err := strconv.Unquote(string(jsonByte))
	if err != nil {
		return dto.ResponseProblem{}, err
	}

	jsonStr = strings.Replace(jsonStr, "```json", "", -1)
	jsonStr = strings.Replace(jsonStr, "```", "", -1)

	var response dto.ResponseProblem
	err = jsoniter.Unmarshal([]byte(jsonStr), &response)
	if err != nil {
		return dto.ResponseProblem{}, err
	}

	return response, nil
}
