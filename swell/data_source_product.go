package swell

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/sweeneyb/swell-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProducts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProductsRead,
		Schema: map[string]*schema.Schema{
			"products": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sku": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceProduct() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProductRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sku": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceProductsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client, err := swell.NewClient()
	if err != nil {
		log.Fatal(err)
		return diag.FromErr(err)
	}
	results, err := client.GetProducts()
	if err != nil {
		log.Fatal(err)
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics
	d.Set("products", flattenProductsData(results))

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func dataSourceProductRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client, err := swell.NewClient()
	if err != nil {
		log.Fatal(err)
		return diag.FromErr(err)
	}
	results, err := client.GetProducts()
	if err != nil {
		log.Fatal(err)
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics
	found := false
	// TODO: once the client library supports a get by ID, swap this out
	for _, p := range results {
		if p.Name == d.Get("name") {
			d.Set("name", p.Name)
			d.Set("sku", p.Sku)
			found = true
		}
	}
	if !found {
		d.Set("product", make([]interface{}, 0))
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenProductsData(products []swell.Product) []interface{} {
	if products != nil {
		ois := make([]interface{}, len(products), len(products))

		for i, product := range products {
			oi := make(map[string]interface{})

			oi["sku"] = product.Sku
			oi["name"] = product.Name

			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)
}
