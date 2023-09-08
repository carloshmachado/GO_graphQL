RUN: go run cmd/server/server.go
open localhost:8080 on Browser
paste:
--------------------
mutation createCategory {
  createCategory(input: {name: "Tecnologia", description: "Cursos de Tecnologia"}) {
    id
    name
    description
  }
}

mutation createCourse {
  createCourse(input: {name: "Full Cycle", description: "The best!", categoryID: "447d4a5a-403a-4b9a-97dd-9eaf440eb547"}) {
	    id
    	name
  }
}

query queryCategories {
  categories {
    id
    name
    description
  }
}

query queryCategoriesWithCourses {
  categories {
    id
    name
    courses {
      id
      name
    }
  }
}

query queryCourses {
  courses {
    id
    name
  }
}

query queryCoursesWithCategory {
  courses {
    id
    name
    description
    category {
      id
      name
      description
    }
  }
}
----------------------