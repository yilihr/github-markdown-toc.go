package main

import "testing"

func Test_IsUrl(t *testing.T) {
	url1 := "https://github.com/ekalinin/envirius/blob/master/README.md"
	if !IsURL(url1) {
		t.Error("This is url: ", url1)
	}

	url2 := "./README.md"
	if IsURL(url2) {
		t.Error("This is not url: ", url2)
	}
}

func Test_GrabTocOneRow(t *testing.T) {
	tocExpected := []string{
		"* [README in another language](#readme-in-another-language)",
	}
	toc := *GrabToc(`
	<h1><a id="user-content-readme-in-another-language" class="anchor" href="#readme-in-another-language" aria-hidden="true"><span class="octicon octicon-link"></span></a>README in another language</h1>
	`, "", 0, 2)
	if toc[0] != tocExpected[0] {
		t.Error("Res :", toc, "\nExpected     :", tocExpected)
	}
}

func Test_GrabTocOneRowWithNewLines(t *testing.T) {
	tocExpected := []string{
		"* [README in another language](#readme-in-another-language)",
	}
	toc := *GrabToc(`
	<h1>
		<a id="user-content-readme-in-another-language" class="anchor" href="#readme-in-another-language" aria-hidden="true">
			<span class="octicon octicon-link"></span>
		</a>
		README in another language
	</h1>
	`, "", 0, 2)
	if toc[0] != tocExpected[0] {
		t.Error("Res :", toc, "\nExpected     :", tocExpected)
	}
}

func Test_GrabTocMultilineOriginGithub(t *testing.T) {

	tocExpected := []string{
		"* [How to add a plugin?](#how-to-add-a-plugin)",
		"  * [Mandatory elements](#mandatory-elements)",
		"    * [plug\\_list\\_versions](#plug_list_versions)",
	}
	toc := *GrabToc(`
<h1><a id="user-content-how-to-add-a-plugin" class="anchor" href="#how-to-add-a-plugin" aria-hidden="true"><span class="octicon octicon-link"></span></a>How to add a plugin?</h1>

<p>All plugins are in the directory
<a href="https://github.com/ekalinin/envirius/tree/master/src/nv-plugins">nv-plugins</a>.
If you need to add support for a new language you should add it as plugin
inside this directory.</p>

<h2><a id="user-content-mandatory-elements" class="anchor" href="#mandatory-elements" aria-hidden="true"><span class="octicon octicon-link"></span></a>Mandatory elements</h2>

<p>If you create a plugin which builds all stuff from source then In a simplest
case you need to implement 2 functions in the plugin's body:</p>

<h3><a id="user-content-plug_list_versions" class="anchor" href="#plug_list_versions" aria-hidden="true"><span class="octicon octicon-link"></span></a>plug_list_versions</h3>

<p>This function should return list of available versions of the plugin.
For example:</p>
	`, "", 0, 2)
	for i := 0; i <= len(tocExpected)-1; i++ {
		if toc[i] != tocExpected[i] {
			t.Error("Res :", toc[i], "\nExpected     :", tocExpected[i])
		}
	}
}

func Test_GrabTocBackquoted(t *testing.T) {
	tocExpected := []string{
		"* [The command foo1](#the-command-foo1)",
		"  * [The command foo2 is better](#the-command-foo2-is-better)",
		"* [The command bar1](#the-command-bar1)",
		"  * [The command bar2 is better](#the-command-bar2-is-better)",
	}

	toc := *GrabToc(`
<h1>
<a id="user-content-the-command-foo1" class="anchor" href="#the-command-foo1" aria-hidden="true"><span class="octicon octicon-link"></span></a>The command <code>foo1</code>
</h1>

<p>Blabla...</p>

<h2>
<a id="user-content-the-command-foo2-is-better" class="anchor" href="#the-command-foo2-is-better" aria-hidden="true"><span class="octicon octicon-link"></span></a>The command <code>foo2</code> is better</h2>

<p>Blabla...</p>

<h1>
<a id="user-content-the-command-bar1" class="anchor" href="#the-command-bar1" aria-hidden="true"><span class="octicon octicon-link"></span></a>The command <code>bar1</code>
</h1>

<p>Blabla...</p>

<h2>
<a id="user-content-the-command-bar2-is-better" class="anchor" href="#the-command-bar2-is-better" aria-hidden="true"><span class="octicon octicon-link"></span></a>The command <code>bar2</code> is better</h2>

<p>Blabla...</p>
	`, "", 0, 2)

	for i := 0; i <= len(tocExpected)-1; i++ {
		if toc[i] != tocExpected[i] {
			t.Error("Res :", toc[i], "\nExpected      :", tocExpected[i])
		}
	}
}

func Test_GrabTocDepth(t *testing.T) {
	tocExpected := []string{
		"* [The command foo1](#the-command-foo1)",
		"* [The command bar1](#the-command-bar1)",
	}

	toc := *GrabToc(`
<h1>
<a id="user-content-the-command-foo1" class="anchor" href="#the-command-foo1" aria-hidden="true"><span class="octicon octicon-link"></span></a>The command <code>foo1</code>
</h1>

<p>Blabla...</p>

<h2>
<a id="user-content-the-command-foo2-is-better" class="anchor" href="#the-command-foo2-is-better" aria-hidden="true"><span class="octicon octicon-link"></span></a>The command <code>foo2</code> is better</h2>

<p>Blabla...</p>

<h1>
<a id="user-content-the-command-bar1" class="anchor" href="#the-command-bar1" aria-hidden="true"><span class="octicon octicon-link"></span></a>The command <code>bar1</code>
</h1>

<p>Blabla...</p>

<h2>
<a id="user-content-the-command-bar2-is-better" class="anchor" href="#the-command-bar2-is-better" aria-hidden="true"><span class="octicon octicon-link"></span></a>The command <code>bar2</code> is better</h2>

<p>Blabla...</p>
	`, "", 1, 2)

	// fmt.Println(toc)

	for i := 0; i <= len(tocExpected)-1; i++ {
		if toc[i] != tocExpected[i] {
			t.Error("Res :", toc[i], "\nExpected      :", tocExpected[i])
		}
	}
}

func Test_GrabTocWithAbspath(t *testing.T) {
	link := "https://github.com/ekalinin/envirius/blob/master/README.md"
	tocExpected := []string{
		"* [README in another language](" + link + "#readme-in-another-language)",
	}
	toc := *GrabToc(`
	<h1><a id="user-content-readme-in-another-language" class="anchor" href="#readme-in-another-language" aria-hidden="true"><span class="octicon octicon-link"></span></a>README in another language</h1>
	`, link, 0, 2)
	if toc[0] != tocExpected[0] {
		t.Error("Res :", toc, "\nExpected     :", tocExpected)
	}
}

func Test_EscapedChars(t *testing.T) {
	tocExpected := []string{
		"* [mod\\_\\*](#mod_)",
	}

	toc := *GrabToc(`
		<h2>
			<a id="user-content-mod_" class="anchor"
			    href="#mod_" aria-hidden="true">
				<span class="octicon octicon-link"></span>
			</a>
			mod_*
		</h2>`, "", 0, 2)

	if toc[0] != tocExpected[0] {
		t.Error("Res :", toc, "\nExpected     :", tocExpected)
	}
}

func Test_CustomSpaceIndentation(t *testing.T) {
	tocExpected := []string{
		"* [Header Level1](#header-level1)",
		"    * [Header Level2](#header-level2)",
		"        * [Header Level3](#header-level3)",
	}

	toc := *GrabToc(`
<h1>
<a id="user-content-the-command-level1" class="anchor" href="#header-level1" aria-hidden="true"><span class="octicon octicon-link"></span></a>Header Level1
</h1>
<h2>
<a id="user-content-the-command-level2" class="anchor" href="#header-level2" aria-hidden="true"><span class="octicon octicon-link"></span></a>Header Level2
</h2>
<h3>
<a id="user-content-the-command-level3" class="anchor" href="#header-level3" aria-hidden="true"><span class="octicon octicon-link"></span></a>Header Level3
</h3>
	`, "", 0, 4) // use 4 spaces indent

	for i := 0; i <= len(tocExpected)-1; i++ {
		if toc[i] != tocExpected[i] {
			t.Error("Res :", toc[i], "\nExpected      :", tocExpected[i])
		}
	}
}

func Test_MinHeaderNumber(t *testing.T) {
	tocExpected := []string{
		"* [foo](#foo)",
		"  * [bar](#bar)",
	}

	toc := *GrabToc(`
		<h3>
			<a id="user-content-" class="anchor" href="#foo" aria-hidden="true">
				<span class="octicon octicon-link"></span>
			</a>
			foo
		</h3>
		<h4>
			<a id="user-content-" class="anchor" href="#bar" aria-hidden="true">
				<span class="octicon octicon-link"></span>
			</a>
			bar
		</h3>
		`, "", 0, 2)

	if toc[0] != tocExpected[0] {
		t.Error("Res :", toc, "\nExpected     :", tocExpected)
	}
}
