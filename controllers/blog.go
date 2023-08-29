package controllers

import (
	"io"
	"log"

	"github.com/devetek/caddy-dataprovider/utils"
)

// Server side rendering model for /blog
func GetDataBlogHome() string {
	apiResp, err := utils.Fetcher("POST", "https://graphql.domain.com", "[{\"operationName\":\"categoryTree\",\"variables\":{\"level\":0,\"source\":\"blog\"},\"query\":\"query categoryTree($level: Int, $source: String) {\\n  categoryTree(level: $level, source: $source) {\\n    data {\\n      status\\n      categories {\\n        id\\n        url\\n        children {\\n          id\\n          url\\n          __typename\\n        }\\n        __typename\\n      }\\n      __typename\\n    }\\n    __typename\\n  }\\n}\\n\"},{\"operationName\":\"GetPosts\",\"variables\":{\"limit\":4,\"filter\":\"highlights\",\"offset\":0,\"source\":\"blog\",\"sort_by\":\"start_datetime\"},\"query\":\"query GetPosts($filter: String!, $limit: Int, $offset: Int, $category_id: Int, $source: String, $sort_by: String) {\\n  cardsArticle(filter: $filter, limit: $limit, offset: $offset, category_id: $category_id, source: $source, sort_by: $sort_by) {\\n    data {\\n      status\\n      cards {\\n        has_more\\n        title\\n        items {\\n          id\\n          title\\n          description\\n          slug\\n          modified_date\\n          publish_time\\n          categories {\\n            url\\n            id\\n            __typename\\n          }\\n          attributes {\\n            read_time\\n            __typename\\n          }\\n          thumbnail {\\n            desktop\\n            mobile\\n            ios\\n            android\\n            __typename\\n          }\\n          __typename\\n        }\\n        total_count\\n        __typename\\n      }\\n      __typename\\n    }\\n    __typename\\n  }\\n}\\n\"},{\"operationName\":\"GetPosts\",\"variables\":{\"limit\":3,\"filter\":\"latest\",\"offset\":0,\"category_id\":315,\"source\":\"blog\",\"sort_by\":\"start_datetime\"},\"query\":\"query GetPosts($filter: String!, $limit: Int, $offset: Int, $category_id: Int, $source: String, $sort_by: String) {\\n  cardsArticle(filter: $filter, limit: $limit, offset: $offset, category_id: $category_id, source: $source, sort_by: $sort_by) {\\n    data {\\n      status\\n      cards {\\n        has_more\\n        title\\n        items {\\n          id\\n          title\\n          description\\n          slug\\n          modified_date\\n          publish_time\\n          categories {\\n            url\\n            id\\n            __typename\\n          }\\n          attributes {\\n            read_time\\n            __typename\\n          }\\n          thumbnail {\\n            desktop\\n            mobile\\n            ios\\n            android\\n            __typename\\n          }\\n          __typename\\n        }\\n        total_count\\n        __typename\\n      }\\n      __typename\\n    }\\n    __typename\\n  }\\n}\\n\"},{\"operationName\":\"GetPosts\",\"variables\":{\"limit\":3,\"filter\":\"latest\",\"offset\":0,\"category_id\":227,\"source\":\"blog\",\"sort_by\":\"start_datetime\"},\"query\":\"query GetPosts($filter: String!, $limit: Int, $offset: Int, $category_id: Int, $source: String, $sort_by: String) {\\n  cardsArticle(filter: $filter, limit: $limit, offset: $offset, category_id: $category_id, source: $source, sort_by: $sort_by) {\\n    data {\\n      status\\n      cards {\\n        has_more\\n        title\\n        items {\\n          id\\n          title\\n          description\\n          slug\\n          modified_date\\n          publish_time\\n          categories {\\n            url\\n            id\\n            __typename\\n          }\\n          attributes {\\n            read_time\\n            __typename\\n          }\\n          thumbnail {\\n            desktop\\n            mobile\\n            ios\\n            android\\n            __typename\\n          }\\n          __typename\\n        }\\n        total_count\\n        __typename\\n      }\\n      __typename\\n    }\\n    __typename\\n  }\\n}\\n\"},{\"operationName\":\"GetPosts\",\"variables\":{\"limit\":3,\"filter\":\"latest\",\"offset\":0,\"category_id\":246,\"source\":\"blog\",\"sort_by\":\"start_datetime\"},\"query\":\"query GetPosts($filter: String!, $limit: Int, $offset: Int, $category_id: Int, $source: String, $sort_by: String) {\\n  cardsArticle(filter: $filter, limit: $limit, offset: $offset, category_id: $category_id, source: $source, sort_by: $sort_by) {\\n    data {\\n      status\\n      cards {\\n        has_more\\n        title\\n        items {\\n          id\\n          title\\n          description\\n          slug\\n          modified_date\\n          publish_time\\n          categories {\\n            url\\n            id\\n            __typename\\n          }\\n          attributes {\\n            read_time\\n            __typename\\n          }\\n          thumbnail {\\n            desktop\\n            mobile\\n            ios\\n            android\\n            __typename\\n          }\\n          __typename\\n        }\\n        total_count\\n        __typename\\n      }\\n      __typename\\n    }\\n    __typename\\n  }\\n}\\n\"}]")
	if err != nil {
		log.Println(err)
		return "[{\"errors\":[{\"message\":\"Request not allowed\",\"extensions\":{\"Code\":406}}]}]"
	}

	bodyBytes, err := io.ReadAll(apiResp.Body)
	if err != nil {
		log.Fatal(err)

		return "[{\"errors\":[{\"message\":\"Request not allowed\",\"extensions\":{\"Code\":406}}]}]"
	}
	bodyString := string(bodyBytes)

	defer apiResp.Body.Close()

	return bodyString
}
