package main

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type UpdateTask struct {
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

type CreateTask struct {
	Title string    `json:"title"`
	Time  int64     `json:"time"`
	Date  time.Time `json:"date"`
}

func main() {
	Initialize()
	app := fiber.New()
	repo := NewTaskRepository(db)

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	app.Get("/tasks", func(c fiber.Ctx) error {
		data, err := repo.GetAllTasks()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.JSON(fiber.Map{
			"data": data,
		})
	})

	app.Get("/tasks/:id", func(c fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		data, err := repo.GetTaskById(uint(id))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data": data,
		})

	})

	app.Patch("/tasks/:id", func(c fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		data := new(UpdateTask)

		if err := c.Bind().Body(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if err := repo.UpdateTask(uint(id), data.Title, data.Date); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusNoContent).Send(nil)
	})
	app.Post("/tasks", func(c fiber.Ctx) error {
		data := new(CreateTask)

		if err := c.Bind().Body(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		task, err := repo.CreateTask(data.Title, data.Time, data.Date)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"data": task,
		})
	})

	app.Delete("/tasks/:id", func(c fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := repo.DeleteTask(uint(id)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusNoContent).Send(nil)
	})
	app.Listen("127.0.0.1:8080")
}
