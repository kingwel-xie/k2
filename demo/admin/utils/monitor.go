package utils

import "os"

func AppMonitor(ret map[string]interface{}) map[string]interface{} {

	app := map[string]interface{}{}

	for _, v := range []struct {
		key string
		env string
		def string
	}{
		{"tenant", "APP_TENANT", "OWL"},
		{"env", "APP_ENV", "dev"},
		{"version", "APP_VERSION", "dev"},
		{"gitCommit", "GIT_COMMIT", "nightly"},
		{"buildTime", "BUILD_TIME", ""},
	} {
		val := getFromEnvOrDefault(v.env, v.def)
		app[v.key] = val
	}

	ret["app"] = app

	return ret
}

func getFromEnvOrDefault(env string, def string) string {
	if val, ok := os.LookupEnv(env); ok {
		return val
	}
	return def
}
