package template

// DO NOT EDIT
// This code is generated.

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/vault/api"
	"github.com/terraform-providers/terraform-provider-vault/util"
)

const nameEndpoint = "/transform/template/{name}"

func NameResource() *schema.Resource {
	fields := map[string]*schema.Schema{
		"path": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Path to backend to configure.",
			StateFunc: func(v interface{}) string {
				return strings.Trim(v.(string), "/")
			},
		},
		"alphabet": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The alphabet to use for this template. This is only used during FPE transformations.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the template.",
			ForceNew:    true,
		},
		"pattern": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The pattern used for matching. Currently, only regular expression pattern is supported.",
		},
		"type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The pattern type to use for match detection. Currently, only regex is supported.",
		},
	}
	return &schema.Resource{
		Create: nameCreateResource,
		Update: nameUpdateResource,
		Read:   nameReadResource,
		Exists: nameResourceExists,
		Delete: nameDeleteResource,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: fields,
	}
}
func nameCreateResource(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)

	path := d.Get("path").(string)

	data := map[string]interface{}{}
	if v, ok := d.GetOkExists("alphabet"); ok {
		data["alphabet"] = v
	}
	if v, ok := d.GetOkExists("name"); ok {
		data["name"] = v
	}
	if v, ok := d.GetOkExists("pattern"); ok {
		data["pattern"] = v
	}
	if v, ok := d.GetOkExists("type"); ok {
		data["type"] = v
	}

	log.Printf("[DEBUG] Writing %q", path)
	_, err := client.Logical().Write(util.ParsePath(path, nameEndpoint, d), data)
	if err != nil {
		return fmt.Errorf("error writing %q: %s", path, err)
	}
	d.SetId(path)
	log.Printf("[DEBUG] Wrote %q", path)
	return nameReadResource(d, meta)
}

func nameReadResource(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	path := d.Id()

	log.Printf("[DEBUG] Reading %q", path)
	resp, err := client.Logical().Read(util.ParsePath(path, nameEndpoint, d))
	if err != nil {
		return fmt.Errorf("error reading %q: %s", path, err)
	}
	log.Printf("[DEBUG] Read %q", path)
	if resp == nil {
		log.Printf("[WARN] %q not found, removing from state", path)
		d.SetId("")
		return nil
	}
	val, ok := resp.Data["alphabet"]
	if !ok {
		continue
	}
	if err := d.Set("alphabet", val); err != nil {
		return fmt.Errorf("error setting state key 'alphabet': %s", err)
	}
	val, ok := resp.Data["name"]
	if !ok {
		continue
	}
	if err := d.Set("name", val); err != nil {
		return fmt.Errorf("error setting state key 'name': %s", err)
	}
	val, ok := resp.Data["pattern"]
	if !ok {
		continue
	}
	if err := d.Set("pattern", val); err != nil {
		return fmt.Errorf("error setting state key 'pattern': %s", err)
	}
	val, ok := resp.Data["type"]
	if !ok {
		continue
	}
	if err := d.Set("type", val); err != nil {
		return fmt.Errorf("error setting state key 'type': %s", err)
	}
	return nil
}

func nameUpdateResource(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	path := d.Id()

	log.Printf("[DEBUG] Updating %q", path)

	data := map[string]interface{}{}
	if d.HasChange("alphabet") {
		data["alphabet"] = d.Get("alphabet")
	}
	if d.HasChange("name") {
		data["name"] = d.Get("name")
	}
	if d.HasChange("pattern") {
		data["pattern"] = d.Get("pattern")
	}
	if d.HasChange("type") {
		data["type"] = d.Get("type")
	}
	defer func() {
		d.SetId(path)
	}()
	_, err := client.Logical().Write(util.ParsePath(path, nameEndpoint, d), data)
	if err != nil {
		return fmt.Errorf("error updating template auth backend role %q: %s", path, err)
	}
	log.Printf("[DEBUG] Updated %q", path)
	return nameReadResource(d, meta)
}

func nameDeleteResource(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	path := d.Id()

	log.Printf("[DEBUG] Deleting %q", path)
	_, err := client.Logical().Delete(path)
	if err != nil && !util.Is404(err) {
		return fmt.Errorf("error deleting %q", path)
	} else if err != nil {
		log.Printf("[DEBUG] %q not found, removing from state", path)
		d.SetId("")
		return nil
	}
	log.Printf("[DEBUG] Deleted template auth backend role %q", path)
	return nil
}

func nameResourceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(*api.Client)

	path := d.Id()
	log.Printf("[DEBUG] Checking if %q exists", path)

	resp, err := client.Logical().Read(path)
	if err != nil {
		return true, fmt.Errorf("error checking if %q exists: %s", path, err)
	}
	log.Printf("[DEBUG] Checked if %q exists", path)
	return resp != nil, nil
}
