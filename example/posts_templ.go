// Code generated by templ DO NOT EDIT.

package main

import "github.com/a-h/templ"
import "context"
import "io"
import "fmt"
import "time"

func headerTemplate() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		_, err = io.WriteString(w, "<head>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<title>")
		if err != nil {
			return err
		}
		var_1 := `Blog`
		_, err = io.WriteString(w, var_1)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</title>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</head>")
		if err != nil {
			return err
		}
		return err
	})
}

func footerTemplate() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		_, err = io.WriteString(w, "<div>")
		if err != nil {
			return err
		}
		var_2 := `&copy; `
		_, err = io.WriteString(w, var_2)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(fmt.Sprintf("%d", time.Now().Year())))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		return err
	})
}

func navTemplate() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		_, err = io.WriteString(w, "<ul>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<li>")
		if err != nil {
			return err
		}
		err = templ.RenderScripts(ctx, w, )
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<a")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " href=\"/\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		var_3 := `Home`
		_, err = io.WriteString(w, var_3)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</a>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</li>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<li>")
		if err != nil {
			return err
		}
		err = templ.RenderScripts(ctx, w, )
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<a")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " href=\"/posts\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		var_4 := `Posts`
		_, err = io.WriteString(w, var_4)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</a>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</li>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</ul>")
		if err != nil {
			return err
		}
		return err
	})
}

func layout(name string, content templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		_, err = io.WriteString(w, "<html>")
		if err != nil {
			return err
		}
		err = headerTemplate().Render(ctx, w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<body>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<h1>")
		if err != nil {
			return err
		}
		var_5 := `Home`
		_, err = io.WriteString(w, var_5)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</h1>")
		if err != nil {
			return err
		}
		err = navTemplate().Render(ctx, w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<main>")
		if err != nil {
			return err
		}
		err = content.Render(ctx, w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</main>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</body>")
		if err != nil {
			return err
		}
		err = footerTemplate().Render(ctx, w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</html>")
		if err != nil {
			return err
		}
		return err
	})
}

func homeTemplate() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		_, err = io.WriteString(w, "<div>")
		if err != nil {
			return err
		}
		var_6 := `Welcome to my website.`
		_, err = io.WriteString(w, var_6)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		return err
	})
}

func home() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		err = layout("Home", homeTemplate()).Render(ctx, w)
		if err != nil {
			return err
		}
		return err
	})
}

func postsTemplate(posts []Post) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		for _, p := range posts {
			err = postTemplate(p).Render(ctx, w)
			if err != nil {
				return err
			}
		}
		return err
	})
}

func postTemplate(post Post) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		_, err = io.WriteString(w, "<div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(post.Name))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, templ.EscapeString(post.Author))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		return err
	})
}

func posts(posts []Post) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		err = layout("Posts", postsTemplate(posts)).Render(ctx, w)
		if err != nil {
			return err
		}
		return err
	})
}

