package rest

import (
	"fmt"
	"strconv"
	"time"

	"drexel.edu/voter-api/pkg/process"
	"drexel.edu/voter-api/pkg/retrieve"
	"github.com/gofiber/fiber/v2"
)

func Handler(port int, processService process.Service, retrievalService retrieve.Service) *fiber.App {
	startTime := time.Now()

	router := fiber.New()

	//GET /voters/health - Returns a "health" record indicating that the voter API is functioning properly and some metadata about the API.  Note the payload can be hard coded, we are mainly looking for a HTTP status code of 200, which means the API is functioning properly.
	router.Get("/voters/health", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK)

		var msg string

		routes := router.GetRoutes()
		for _, r := range routes {
			msg = fmt.Sprintf("%v\r\n method: %s \r\n path: %s\r\n", msg, r.Method, r.Path)
		}

		msg = fmt.Sprintf("OK!!!  Uptime: %v \r\nRoutes:\r\n%s", time.Since(startTime), msg)

		return c.SendString(msg)
	})

	//GET /voters - Get all voter resources including all voter history for each voter (note we will discuss the concept of "paging" later, for now you can ignore)
	router.Get("/voters", func(c *fiber.Ctx) error {

		votersDTO, err := retrievalService.GetAllVoters()
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return err
		}

		var voters []Voter

		for _, voter := range votersDTO {
			voters = append(voters, convertVoterToMuteable(voter))
		}

		c.Status(fiber.StatusOK)
		return c.JSON(voters)
	})

	//GET&POST /voters/:id - Get a single voter resource with voterID=:id including their entire voting history.  POST version adds one to the "database"
	router.Get("/voters/:id", func(c *fiber.Ctx) error {

		voterId, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		voterDTO, err := retrievalService.GetSingleVoter(voterId)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return err
		}

		c.Status(fiber.StatusOK)
		return c.JSON(convertVoterToMuteable(voterDTO))
	})

	router.Post("/voters/:id", func(c *fiber.Ctx) error {

		var voter Voter

		voterId, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		if err := c.BodyParser(&voter); err != nil {
			return err
		}

		voterDTO := process.NewVoterDTO(
			voterId,
			voter.Name,
			voter.Email,
		)

		err = processService.CreateVoter(voterDTO)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return err
		}

		c.Status(fiber.StatusCreated)

		return c.SendString("Voter registration successful.")
	})

	//GET /voters/:id/polls - Gets the JUST the voter history for the voter with VoterID = :id
	router.Get("/voters/:id/polls", func(c *fiber.Ctx) error {
		var voter []retrieve.VoterHistoryDTO

		c.Status(fiber.StatusInternalServerError)

		voterId, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		voter, err = retrievalService.GetVoterHistory(voterId)
		if err != nil {
			return err
		}

		var history []VoterHistory

		for _, poll := range voter {
			history = append(history, convertHistoryToMuteable(poll))
		}

		c.Status(fiber.StatusOK)
		return c.JSON(history)

	})

	//GET&POST /voters/:id/polls/:pollid - Gets JUST the single voter poll data with PollID = :id and VoterID = :id.  POST version adds one to the "database"

	router.Get("/voters/:voterId/polls/:pollId", func(c *fiber.Ctx) error {
		var voter retrieve.VoterHistoryDTO

		c.Status(fiber.StatusInternalServerError)

		voterId, err := strconv.Atoi(c.Params("voterId"))
		if err != nil {
			return err
		}

		pollId, err := strconv.Atoi(c.Params("pollId"))
		if err != nil {
			return err
		}

		voter, err = retrievalService.GetSingleEvent(voterId, pollId)
		if err != nil {
			return err
		}

		c.Status(fiber.StatusOK)
		return c.JSON(convertHistoryToMuteable(voter))

	})

	router.Post("/voters/:voterId/polls/:pollId", func(c *fiber.Ctx) error {
		var voterHistory VoterHistory

		c.Status(fiber.StatusInternalServerError)

		voterId, err := strconv.Atoi(c.Params("voterId"))
		if err != nil {
			return err
		}

		pollId, err := strconv.Atoi(c.Params("pollId"))
		if err != nil {
			return err
		}

		if err := c.BodyParser(&voterHistory); err != nil {
			return err
		}

		voteDate, err := time.Parse(time.RFC3339, voterHistory.VoteDate)
		if err != nil {
			return err
		}

		historyDTO := process.NewVoterHistoryDTO(
			pollId,
			voterId,
			voteDate,
		)

		err = processService.CreateVoterHistory(voterId, pollId, historyDTO)
		if err != nil {
			return err
		}

		c.Status(fiber.StatusCreated)

		return c.SendString("CREATED")

	})

	router.Put("/voters/:id", func(c *fiber.Ctx) error {

		var voter Voter

		voterId, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		if err := c.BodyParser(&voter); err != nil {
			return err
		}

		voterDTO := process.NewVoterDTO(
			voterId,
			voter.Name,
			voter.Email,
		)

		err = processService.UpdateVoterInfo(voterDTO)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return err
		}

		c.Status(fiber.StatusOK)

		return c.SendString("Voter update successful.")
	})

	router.Put("/voters/:voterId/polls/:pollId", func(c *fiber.Ctx) error {
		var voterHistory VoterHistory

		c.Status(fiber.StatusInternalServerError)

		voterId, err := strconv.Atoi(c.Params("voterId"))
		if err != nil {
			return err
		}

		pollId, err := strconv.Atoi(c.Params("pollId"))
		if err != nil {
			return err
		}

		if err := c.BodyParser(&voterHistory); err != nil {
			return err
		}

		voteDate, err := time.Parse(time.RFC3339, voterHistory.VoteDate)
		if err != nil {
			return err
		}

		historyDTO := process.NewVoterHistoryDTO(
			pollId,
			voterId,
			voteDate,
		)

		err = processService.UpdateVoterHistoryInfo(voterId, pollId, historyDTO)
		if err != nil {
			return err
		}

		c.Status(fiber.StatusOK)

		return c.SendString("The voter history was successfully updated.")

	})

	router.Delete("/voters/:id", func(c *fiber.Ctx) error {

		voterId, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		err = processService.DeleteSingleVoter(voterId)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return err
		}

		c.Status(fiber.StatusOK)

		return c.SendString("Voter was removed.")
	})

	router.Delete("/voters/:voterId/polls/:pollId", func(c *fiber.Ctx) error {

		c.Status(fiber.StatusInternalServerError)

		voterId, err := strconv.Atoi(c.Params("voterId"))
		if err != nil {
			return err
		}

		pollId, err := strconv.Atoi(c.Params("pollId"))
		if err != nil {
			return err
		}

		err = processService.DeleteSingleVoterPoll(voterId, pollId)
		if err != nil {
			return err
		}

		c.Status(fiber.StatusOK)

		return c.SendString("The voter history was successfully deleted.")

	})

	return router
}

func convertVoterToMuteable(voterDTO retrieve.VoterDTO) Voter {
	voter := Voter{
		Id:       voterDTO.GetId(),
		Name:     voterDTO.GetName(),
		Email:    voterDTO.GetEmail(),
		Created:  voterDTO.GetCreated().Format(time.RFC3339),
		Modified: voterDTO.GetModified().Format(time.RFC3339),
	}

	for _, item := range voterDTO.GetHistory() {
		voter.VoterHistory = append(voter.VoterHistory, VoterHistory{
			PollId:   item.GetPollID(),
			VoteId:   item.GetVoteID(),
			VoteDate: item.GetVoteDate().Format(time.RFC3339),
			Created:  item.GetCreated().Format(time.RFC3339),
			Modified: item.GetModified().Format(time.RFC3339),
		})
	}

	return voter
}

func convertHistoryToMuteable(historyDTO retrieve.VoterHistoryDTO) VoterHistory {
	history := VoterHistory{
		PollId:   historyDTO.GetPollID(),
		VoteId:   historyDTO.GetVoteID(),
		VoteDate: historyDTO.GetVoteDate().Format(time.RFC3339),
		Created:  historyDTO.GetCreated().Format(time.RFC3339),
		Modified: historyDTO.GetModified().Format(time.RFC3339),
	}
	return history
}
