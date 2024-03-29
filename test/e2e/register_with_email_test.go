package e2e

import (
	"context"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/octofoxio/sparkle"
	commonv1 "github.com/octofoxio/sparkle/pkg/common/v1"
	entitiesv1 "github.com/octofoxio/sparkle/pkg/entities/v1"
	"github.com/octofoxio/sparkle/pkg/rand"
	svcsv1 "github.com/octofoxio/sparkle/pkg/svcs/v1"
	"github.com/octofoxio/sparkle/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func Test_E2E_RegisterWithEmail(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}
	testutils.NewSuite(t, func(t *testing.T, database sparkle.Database, clients *testutils.SuiteClients) {
		ctx := context.Background()
		registerInputData := rand.RegisterWithEmailInput()
		registerInputData.CallbackURL = commonv1.NotNullString("https://www.google.com")
		output, err := clients.Sparkle.Register(ctx, &svcsv1.RegisterInput{
			RegisterInputData: &svcsv1.RegisterInput_Email{
				Email: registerInputData,
			},
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, output)

		sessionCollections := database.Collection(sparkle.SessionCollectionName)
		var session entitiesv1.Session
		err = sessionCollections.FindOne(ctx, &entitiesv1.Session{UserID: output.GetResult().GetID()}, &session)
		assert.NoError(t, err)
		assert.Equal(t, session.GetUserID(), output.GetResult().GetID())
		assert.NotNil(t, session.GetAccessToken())

		t.Run("access token is valid after register and confirm email", func(t *testing.T) {
			validateOutput, err := clients.Sparkle.ValidateAccessToken(context.Background(), &svcsv1.ValidateAccessTokenInput{
				AccessToken: session.GetAccessToken(),
			})
			assert.NoError(t, err)
			assert.EqualValues(t, validateOutput.GetResult().GetIsValid(), true)
		})

		now := time.Now()
		guard := monkey.Patch(time.Now, func() time.Time {
			// because in setup_test.go
			// we use accessTokenLifetime = 1 minute
			return now.Add(time.Hour * 3600)
		})
		t.Run("access token should invalid after not use for a period of time", func(t *testing.T) {
			validateOutput, err := clients.Sparkle.ValidateAccessToken(context.Background(), &svcsv1.ValidateAccessTokenInput{
				AccessToken: session.GetAccessToken(),
			})
			assert.NoError(t, err)
			assert.EqualValues(t, validateOutput.GetResult().GetIsValid(), false)
		})
		defer guard.Restore()
	})

}
