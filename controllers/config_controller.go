package controllers

import "fmt"

// ConfigController estructura que maneja la configuración
type ConfigController struct {
	config map[string]string
}

// NewConfigController crea una nueva instancia de ConfigController
func NewConfigController() *ConfigController {
	return &ConfigController{
		config: make(map[string]string),
	}
}

// Set establece un valor en la configuración
func (c *ConfigController) Set(key string, value string) {
	c.config[key] = value
}

// Get obtiene un valor de la configuración
func (c *ConfigController) Get(key string) (string, error) {
	value, exists := c.config[key]
	if !exists {
		return "", fmt.Errorf("config key '%s' not found", key)
	}
	return value, nil
}

// Delete elimina un valor de la configuración
func (c *ConfigController) Delete(key string) {
	delete(c.config, key)
}

// List devuelve todas las configuraciones
func (c *ConfigController) List() map[string]string {
	return c.config
}
