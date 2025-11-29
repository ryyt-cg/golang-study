


/*
log.Println("Loading config...")
viper.SetConfigName("application")
viper.SetConfigType("yaml")
viper.AddConfigPath(".")
err := viper.ReadInConfig()
if err != nil {
panic("Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
}
log.Println("Config loaded successfully...")
log.Println("Getting environment variables...")
for _, k := range viper.AllKeys() {
value := viper.GetString(k)
if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value,"${"), "}")))
}
}

func getEnvOrPanic(env string) string {
res := os.Getenv(env)
if len(res) == 0 {
panic("Mandatory env variable not found:" + env)
}
return res
}
*/
