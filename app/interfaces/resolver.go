package interfaces

import (
	"github.com/x893675/graphql-sample/app/domain/repository/answer"
	"github.com/x893675/graphql-sample/app/domain/repository/question"
	"github.com/x893675/graphql-sample/app/domain/repository/question_option"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AnsService            answer.AnsService
	QuestionService       question.QuesService
	QuestionOptionService question_option.OptService
}
