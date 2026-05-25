package routes

import (
	"workhub/internal/handlers"
	"workhub/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	authHandler :=
		handlers.NewAuthHandler()

	testHandler :=
		handlers.NewTestHandler()

	companyHandler :=
		handlers.NewCompanyHandler()

	jobHandler :=
		handlers.NewJobHandler()

	applicationHandler :=
		handlers.NewApplicationHandler()

	reportHandler :=
		handlers.NewReportHandler()

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST(
				"/register",
				authHandler.Register,
			)

			auth.POST(
				"/login",
				authHandler.Login,
			)
		}

		protected := api.Group("/")
		protected.Use(
			middleware.AuthMiddleware(),
		)

		{
			protected.GET(
				"/profile",
				testHandler.Profile,
			)
		}

		company := api.Group("/companies")

		company.Use(
			middleware.AuthMiddleware(),
			middleware.RoleMiddleware(
				"employer",
			),
		)

		{
			company.POST(
				"",
				companyHandler.CreateCompany,
			)

			company.GET(
				"/me",
				companyHandler.GetMyCompany,
			)

			company.PUT(
				"",
				companyHandler.UpdateCompany,
			)
		}

		jobs := api.Group("/jobs")
		{
			jobs.GET(
				"",
				jobHandler.GetAllJobs,
			)

			jobs.GET(
				"/:id",
				jobHandler.GetJobByID,
			)

			jobs.POST(
				"",
				middleware.AuthMiddleware(),
				middleware.RoleMiddleware(
					"employer",
				),
				jobHandler.CreateJob,
			)

			jobs.POST(
				"/:id/apply",
				middleware.AuthMiddleware(),
				middleware.RoleMiddleware(
					"jobseeker",
				),
				applicationHandler.ApplyJob,
			)

			jobs.GET(
				"/:id/applications",
				middleware.AuthMiddleware(),
				middleware.RoleMiddleware(
					"employer",
				),
				applicationHandler.
					GetJobApplications,
			)
		}

		applications :=
			api.Group("/applications")

		applications.Use(
			middleware.AuthMiddleware(),
		)

		{

			applications.GET(
				"/me",
				middleware.RoleMiddleware(
					"jobseeker",
				),
				applicationHandler.
					GetMyApplications,
			)

			applications.PATCH(
				"/:id",
				middleware.RoleMiddleware(
					"employer",
				),
				applicationHandler.
					UpdateApplicationStatus,
			)
		}

		admin := api.Group("/admin")

		admin.Use(
			middleware.AuthMiddleware(),
			middleware.RoleMiddleware(
				"admin",
			),
		)

		{
			admin.GET(
				"/test",
				func(c *gin.Context) {
					c.JSON(200, gin.H{
						"message": "admin access granted",
					})
				},
			)

			admin.GET(
				"/companies/pending",
				companyHandler.
					GetPendingCompanies,
			)

			admin.PATCH(
				"/companies/:id/approve",
				companyHandler.
					ApproveCompany,
			)

			admin.PATCH(
				"/companies/:id/reject",
				companyHandler.
					RejectCompany,
			)

			admin.GET(
				"/reports",
				reportHandler.
					GetDashboardReport,
			)
		}
	}
}
