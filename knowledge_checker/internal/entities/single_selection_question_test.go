package entities_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/parta4ok/kvs/knowledge_checker/internal/entities"
)

func TestNewSingleSelectionQuestion(t *testing.T) {
	t.Parallel()

	id := uint64(time.Now().UnixNano())
	topic := "test topic"
	subject := "what is your name?"
	variants := []string{"John Doe", "Jack Spearrow", "Jonn Leanon", "John Travolta"}
	correctAnswer := "John Doe"

	question := entities.NewSingleSelectionQuestion(id, topic, subject, variants, correctAnswer)
	require.NotNil(t, question)

	notCorrectAns, err := entities.NewUserAnswer(question.ID(), []string{"Jack Spearrow"})
	require.NoError(t, err)

	correctAns, err := entities.NewUserAnswer(question.ID(), []string{"John Doe"})
	require.NoError(t, err)

	res := question.IsAnswerCorrect(notCorrectAns)
	require.False(t, res)

	res = question.IsAnswerCorrect(correctAns)
	require.True(t, res)

	require.Equal(t, id, question.ID())
	require.Equal(t, topic, question.Topic())
	require.Equal(t, subject, question.Subject())
	require.Equal(t, variants, question.Variants())
	require.Equal(t, entities.SingleSelection, question.Type())
}
