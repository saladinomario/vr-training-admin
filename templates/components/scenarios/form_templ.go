// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
// templates/components/scenarios/form.templ

package scenarios

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fmt"

func ScenarioForm(scenario *Scenario, isEdit bool) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"card bg-base-100 shadow-xl\"><div class=\"card-body\"><h2 class=\"card-title\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if isEdit {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "Edit Scenario: ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(scenario.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/scenarios/form.templ`, Line: 11, Col: 49}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "Create New Scenario")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</h2><form")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if isEdit {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, " hx-put=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs("/scenarios/" + scenario.ID)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/scenarios/form.templ`, Line: 19, Col: 55}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, " hx-post=\"/scenarios\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, " hx-target=\"#main-content\" hx-swap=\"outerHTML\" class=\"space-y-6\"><!-- Basic Information Section --><div class=\"space-y-4\"><h3 class=\"text-lg font-medium\">Basic Information</h3><div class=\"form-control w-full\"><label class=\"label\"><span class=\"label-text\">Scenario Name</span></label> <input type=\"text\" name=\"name\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(scenario.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/scenarios/form.templ`, Line: 38, Col: 48}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\" placeholder=\"Enter scenario name\" class=\"input input-bordered w-full\" required></div><div class=\"form-control w-full\"><label class=\"label\"><span class=\"label-text\">Description</span></label> <textarea name=\"description\" placeholder=\"Describe the scenario\" class=\"textarea textarea-bordered h-24\" required>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(scenario.Description)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/scenarios/form.templ`, Line: 54, Col: 46}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</textarea></div><div class=\"grid grid-cols-1 md:grid-cols-3 gap-4\"><div class=\"form-control\"><label class=\"label\"><span class=\"label-text\">Category</span></label> <select name=\"category\" class=\"select select-bordered w-full\"><option value=\"\" disabled")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Category == "" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, ">Select category</option> <option value=\"Sales\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Category == "Sales" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, ">Sales</option> <option value=\"Customer Service\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Category == "Customer Service" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, ">Customer Service</option> <option value=\"Leadership\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Category == "Leadership" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, ">Leadership</option> <option value=\"Conflict Resolution\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Category == "Conflict Resolution" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, ">Conflict Resolution</option> <option value=\"Negotiation\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Category == "Negotiation" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, ">Negotiation</option></select></div><div class=\"form-control\"><label class=\"label\"><span class=\"label-text\">Difficulty Level</span></label> <select name=\"difficulty\" class=\"select select-bordered w-full\"><option value=\"1\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Difficulty == 1 {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, ">Beginner</option> <option value=\"2\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Difficulty == 2 {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, ">Easy</option> <option value=\"3\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Difficulty == 3 {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, ">Intermediate</option> <option value=\"4\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Difficulty == 4 {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 29, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 30, ">Advanced</option> <option value=\"5\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Difficulty == 5 {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 31, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 32, ">Expert</option></select></div><div class=\"form-control\"><label class=\"label\"><span class=\"label-text\">Duration (minutes)</span></label> <input type=\"number\" name=\"duration\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprint(scenario.Duration))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/scenarios/form.templ`, Line: 92, Col: 68}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 33, "\" min=\"5\" max=\"120\" step=\"5\" class=\"input input-bordered w-full\"></div></div></div><!-- Environment Settings --><div class=\"space-y-4\"><h3 class=\"text-lg font-medium\">Environment Settings</h3><div class=\"form-control\"><label class=\"label\"><span class=\"label-text\">VR Scene</span></label> <select name=\"scene\" class=\"select select-bordered w-full\"><option value=\"office\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Scene == "office" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 34, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 35, ">Office</option> <option value=\"meeting_room\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Scene == "meeting_room" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 36, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 37, ">Meeting Room</option> <option value=\"retail_store\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Scene == "retail_store" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 38, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 39, ">Retail Store</option> <option value=\"cafe\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Scene == "cafe" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 40, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 41, ">Café</option> <option value=\"outdoor\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.Scene == "outdoor" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 42, " selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 43, ">Outdoor</option></select></div><div class=\"form-control\"><label class=\"label cursor-pointer\"><span class=\"label-text\">Enable Background Noise</span> <input type=\"checkbox\" name=\"background_noise\" class=\"toggle toggle-primary\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if scenario.BackgroundNoise {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 44, " checked")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 45, "></label></div></div><!-- Success Criteria --><div class=\"space-y-4\"><h3 class=\"text-lg font-medium\">Success Criteria</h3><div class=\"form-control\"><label class=\"label\"><span class=\"label-text\">Success Definition</span></label> <textarea name=\"success_criteria\" placeholder=\"Define what constitutes success in this scenario\" class=\"textarea textarea-bordered h-24\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(scenario.SuccessCriteria)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/scenarios/form.templ`, Line: 146, Col: 50}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 46, "</textarea></div><div class=\"form-control w-full\"><label class=\"label\"><span class=\"label-text\">Required Keywords</span> <span class=\"label-text-alt\">Comma separated</span></label> <input type=\"text\" name=\"keywords\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(scenario.Keywords)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/scenarios/form.templ`, Line: 157, Col: 52}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 47, "\" placeholder=\"Enter keywords that should be used\" class=\"input input-bordered w-full\"></div></div><div class=\"card-actions justify-end\"><a href=\"/scenarios\" class=\"btn btn-ghost\">Cancel</a> <button type=\"submit\" class=\"btn btn-primary\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if isEdit {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 48, "Save Changes")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 49, "Create Scenario")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 50, "</button></div></form></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
