// Code generated by templ@v0.2.186 DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

// GoExpression
import "fmt"

func table(t Data) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = new(bytes.Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
// RawElement
		_, err = templBuffer.WriteString("<style")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" type=\"text/css\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
// Text
var_2 := `
		.jsontable-table {
			border-collapse: collapse;
			margin: 25px 0;
			font-size: 0.9em;
			font-family: sans-serif;
			min-width: 400px;
			border: solid 1px #dddddd;
		}
		th.jsontable-pk {
			background-color: #D6EAF8;
			text-align: center;
			padding: 10px;
			font-weight: bolder;
			border: solid 1px #dddddd;
		}
		td.jsontable-pk {
			text-align: center;
			padding: 10px;
			font-weight: bolder;
			border: solid 1px #dddddd;
		}
		th.jsontable-sk {
			background-color: #EBF5FB;
			text-align: center;
			padding: 10px;
			font-weight: bolder;
			border: solid 1px #dddddd;
		}
		td.jsontable-sk {
			text-align: center;
			padding: 10px;
			font-weight: bolder;
			border: solid 1px #dddddd;
		}
		th.jsontable-key {
			background-color: #E9F7EF;
			text-align: center;
			padding: 10px;
			font-weight: bolder;
			border: solid 1px #dddddd;
		}
		td.jsontable-key {
			text-align: center;
			padding: 10px;
			border: solid 1px #dddddd;
		}
		th.jsontable-attr {
			background-color: #eeeeee;
			text-align: center;
			padding: 10px;
			font-weight: bolder;
			border: solid 1px #dddddd;
		}
		td.jsontable-attr {
			text-align: center;
			padding: 10px;
			border: solid 1px #dddddd;
		}
		tr.jsontable-even {
			background-color: #ffffff;
		}
		tr.jsontable-odd {
			background-color: #eeeeee;
		}
	`
_, err = templBuffer.WriteString(var_2)
if err != nil {
	return err
}
		_, err = templBuffer.WriteString("</style>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<table")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"jsontable-table\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<tr>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<th")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"jsontable-pk\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// StringExpression
		_, err = templBuffer.WriteString(templ.EscapeString(t.PK))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th>")
		if err != nil {
			return err
		}
		// If
		if t.SK != "" {
			// Element (standard)
			_, err = templBuffer.WriteString("<th")
			if err != nil {
				return err
			}
			// Element Attributes
			_, err = templBuffer.WriteString(" class=\"jsontable-sk\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(">")
			if err != nil {
				return err
			}
			// StringExpression
			_, err = templBuffer.WriteString(templ.EscapeString(t.SK))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</th>")
			if err != nil {
				return err
			}
		}
		// For
		for _, th := range t.Attributes {
			// Element (standard)
			_, err = templBuffer.WriteString("<th")
			if err != nil {
				return err
			}
			// Element Attributes
			_, err = templBuffer.WriteString(" class=\"jsontable-key\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(">")
			if err != nil {
				return err
			}
			// StringExpression
			_, err = templBuffer.WriteString(templ.EscapeString(th))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</th>")
			if err != nil {
				return err
			}
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<th")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" class=\"jsontable-attr\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" colspan=")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(fmt.Sprintf("%d", t.MaxColCount)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Text
		var_3 := `Attributes`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</tr>")
		if err != nil {
			return err
		}
		// For
		for partitionIndex, p := range t.Partitions {
			// For
			for i, r := range p.Rows {
				// Element (standard)
				// Element CSS
				var var_4 templ.CSSClasses = getRowClass(partitionIndex)
				err = templ.RenderCSSItems(ctx, templBuffer, var_4...)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("<tr")
				if err != nil {
					return err
				}
				// Element Attributes
				_, err = templBuffer.WriteString(" class=")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("\"")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(templ.EscapeString(var_4.String()))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("\"")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(">")
				if err != nil {
					return err
				}
				// If
				if i == 0 {
					// Element (standard)
					_, err = templBuffer.WriteString("<td")
					if err != nil {
						return err
					}
					// Element Attributes
					_, err = templBuffer.WriteString(" class=\"jsontable-pk\"")
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString(" rowspan=")
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString("\"")
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString(templ.EscapeString(fmt.Sprintf("%d", len(p.Rows))))
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString("\"")
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString(">")
					if err != nil {
						return err
					}
					// StringExpression
					_, err = templBuffer.WriteString(templ.EscapeString(p.PK))
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString("</td>")
					if err != nil {
						return err
					}
				}
				// If
				if t.SK != "" {
					// Element (standard)
					_, err = templBuffer.WriteString("<td")
					if err != nil {
						return err
					}
					// Element Attributes
					_, err = templBuffer.WriteString(" class=\"jsontable-sk\"")
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString(">")
					if err != nil {
						return err
					}
					// StringExpression
					_, err = templBuffer.WriteString(templ.EscapeString(getValueOrEmptyString(t.SK, r)))
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString("</td>")
					if err != nil {
						return err
					}
				}
				// For
				for _, attr := range t.Attributes {
					// Element (standard)
					_, err = templBuffer.WriteString("<td")
					if err != nil {
						return err
					}
					// Element Attributes
					_, err = templBuffer.WriteString(" class=\"jsontable-key\"")
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString(">")
					if err != nil {
						return err
					}
					// StringExpression
					_, err = templBuffer.WriteString(templ.EscapeString(getValueOrEmptyString(attr, r)))
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString("</td>")
					if err != nil {
						return err
					}
				}
				// For
				for _, v := range r {
					// If
					if !t.IsAttribute(v.Key) {
						// Element (standard)
						_, err = templBuffer.WriteString("<td")
						if err != nil {
							return err
						}
						// Element Attributes
						_, err = templBuffer.WriteString(" class=\"jsontable-attr\"")
						if err != nil {
							return err
						}
						_, err = templBuffer.WriteString(">")
						if err != nil {
							return err
						}
						// StringExpression
						_, err = templBuffer.WriteString(templ.EscapeString(v.String()))
						if err != nil {
							return err
						}
						_, err = templBuffer.WriteString("</td>")
						if err != nil {
							return err
						}
					}
				}
				// If
				if t.GetAttributeCount(r) < t.MaxColCount {
					// Element (standard)
					_, err = templBuffer.WriteString("<td")
					if err != nil {
						return err
					}
					// Element Attributes
					_, err = templBuffer.WriteString(" class=\"jsontable-attr\"")
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString(" colspan=")
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString("\"")
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString(templ.EscapeString(fmt.Sprintf("%d", t.MaxColCount)))
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString("\"")
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString(">")
					if err != nil {
						return err
					}
					_, err = templBuffer.WriteString("</td>")
					if err != nil {
						return err
					}
				}
				_, err = templBuffer.WriteString("</tr>")
				if err != nil {
					return err
				}
			}
		}
		_, err = templBuffer.WriteString("</table>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
