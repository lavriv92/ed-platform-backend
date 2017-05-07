package routes

import (
	"net/http"

	"app/auth"
	"app/courses"
	"app/users"
)

const (
	API_VERSION = "1"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var authRoutes = Routes{
	Route{
		"Token",
		"POST",
		"/token",
		auth.AuthTokenHandler,
	},
	Route{
		"Users",
		"POST",
		"/users",
		users.UsersCreateHandler,
	},
}

var apiRoutes = Routes{
	Route{
		"Courses",
		"GET",
		"/courses",
		courses.CoursesIndexHandler,
	},
	Route{
		"CoursesCreate",
		"POST",
		"/courses",
		courses.CoursesCreateHandler,
	},
	Route{
		"Course",
		"GET",
		"/courses/{id}",
		courses.CourseGetHandler,
	},
	Route{
		"CourseLessons",
		"GET",
		"/courses/{id}/lessons",
		courses.CourseLessonsHandler,
	},
}
