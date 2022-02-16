package swell

import (
	"context"
	"log"
	"time"

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
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCategoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// TODO this should be from m
	client, _ := swell.NewClient()

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
	client, _ := swell.NewClient()

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

	// TODO this should be from m
	client, _ := swell.NewClient()

	id := d.Id()
	var cat = swell.Category{
		Id: id,
	}
	var anyChanges = false
	if d.HasChange("name") {
		cat.Name = d.Get("name").(string)
		anyChanges = true
	}

	// BUG This never seems to go to false
	if d.HasChange("active") {
		cat.Active = d.Get("active").(bool)
		anyChanges = true
	}
	if d.HasChange("description") {
		cat.Description = d.Get("description").(string)
		anyChanges = true
	}
	if anyChanges {
		_, err := client.UpdateCategory(cat)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceCategoryRead(ctx, d, m)
}

func resourceCategoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// TODO this should be from m
	client, _ := swell.NewClient()

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	var cat = swell.Category{
		Id: id,
	}
	err := client.DeleteCategory(cat)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
