package swell

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sweeneyb/swell-go"
)

func resourceCategory() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCategoryCreate,
		ReadContext:   resourceCategoryRead,
		UpdateContext: resourceCategoryUpdate,
		DeleteContext: resourceCategoryDelete,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"active": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCategoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// TODO this should be from m
	client, err := swell.NewClient()

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var category = swell.Category{
		Name:        d.Get("name").(string),
		Active:      d.Get("active").(bool),
		Description: d.Get("description").(string),
	}

	result, err := client.CreateCategory(category)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(result.Id)
	resourceCategoryRead(ctx, d, m)
	return diags
}

func resourceCategoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// TODO this should be from m
	client, err := swell.NewClient()

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	result, err := client.GetCategory(id)
	if err != nil {
		log.Fatal(err)
		return diag.FromErr(err)
	}
	d.Set("name", result.Name)
	d.Set("active", result.Active)
	d.Set("description", result.Description)
	d.Set("id", result.Id)
	d.SetId(result.Id)
	return diags
}

func resourceCategoryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCategoryRead(ctx, d, m)
}

func resourceCategoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
