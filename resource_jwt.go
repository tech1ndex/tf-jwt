package main

import (
	"github.com/golang-jwt/jwt"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceJwtToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceJwtCreate,
		Read:   resourceJwtRead,
		Delete: resourceJwtDelete,
		Schema: map[string]*schema.Schema{
			"secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "JWT secret key",
			},
			"claims": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "JWT claims",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"token": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Generated JWT token",
			},
		},
	}
}

func resourceJwtCreate(d *schema.ResourceData, m interface{}) error {
	secret := d.Get("secret").(string)
	claims := d.Get("claims").(map[string]interface{})

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claimsMap := make(map[string]interface{})
	for key, value := range claims {
		claimsMap[key] = value
	}
	token.Claims = claimsMap

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	d.Set("token", tokenString)
	d.SetId(tokenString)
	return nil
}

func resourceJwtRead(d *schema.ResourceData, m interface{}) error {
	// Implement read logic if needed
	return nil
}

func resourceJwtDelete(d *schema.ResourceData, m interface{}) error {
	// Implement delete logic if needed
	return nil
}
