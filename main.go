package main

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type UpdateTask struct {
	Title string    `json:"title"`
	Time  time.Time `json:"date"`
}

type CreateTask struct {
	Title    string    `json:"title"`
	DeadLine time.Time `json:"deadline"`
}

func main() {
	Initialize()
	app := fiber.New()
	repo := NewTaskRepository(db)

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://127.0.0.1:5500"},
		AllowMethods: []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH"},
	}))

	app.Get("/tasks", func(c fiber.Ctx) error {
		data, err := repo.GetAllTasks()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
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
		if err := repo.UpdateTask(uint(id), data.Title); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusNoContent).Send(nil)
	})

	app.Post("/tasks", func(c fiber.Ctx) error {
		var data CreateTask

		if err := c.Bind().Body(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		loc, _ := time.LoadLocation("Europe/Moscow")
		data.DeadLine = data.DeadLine.In(loc)

		task, err := repo.CreateTask(data.Title, data.DeadLine)
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
