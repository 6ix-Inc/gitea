	"code.gitea.io/gitea/modules/highlight"
		t.Errorf("Did not receive expected results:\nExpected: %s\nActual:   %s", s1, s2)
	assertEqual(t, "<span class=\"nx\">r</span><span class=\"p\">.</span><span class=\"nf\">WrapperRenderer</span><span class=\"p\">(</span><span class=\"nx\">w</span><span class=\"p\">,</span> <span class=\"removed-code\"><span class=\"nx\">language</span><span class=\"p\">,</span> <span class=\"kc\">true</span><span class=\"p\">,</span> <span class=\"nx\">attrs</span></span><span class=\"p\">,</span> <span class=\"kc\">false</span><span class=\"p\">)</span>", diffToHTML("", []dmp.Diff{
	assertEqual(t, "<span class=\"added-code\">language</span><span class=\"p\">,</span> <span class=\"kc\">true</span><span class=\"p\">,</span> <span class=\"nx\">attrs</span></span><span class=\"p\">,</span> <span class=\"kc\">false</span><span class=\"p\">)</span>", diffToHTML("", []dmp.Diff{
	assertEqual(t, "<span class=\"k\">print</span><span class=\"added-code\"><span class=\"p\">(</span></span><span class=\"sa\"></span><span class=\"s2\">&#34;</span><span class=\"s2\">// </span><span class=\"s2\">&#34;</span><span class=\"p\">,</span> <span class=\"n\">sys</span><span class=\"o\">.</span><span class=\"n\">argv</span><span class=\"added-code\"><span class=\"p\">)</span></span>", diffToHTML("", []dmp.Diff{
	assertEqual(t, "sh <span class=\"added-code\">&#39;useradd -u $(stat -c &#34;%u&#34; .gitignore) jenkins&#39;</span>", diffToHTML("", []dmp.Diff{
	assertEqual(t, "<span class=\"x\">							&lt;h<span class=\"added-code\">4 class=&#34;release-list-title df ac&#34;</span>&gt;</span>", diffToHTML("", []dmp.Diff{
		{
			name: "ambiguous deleted",
			gitdiff: `diff --git a/b b/b b/b b/b
deleted file mode 100644
index 92e798b..0000000
--- a/b b/b` + "\t" + `
+++ /dev/null
@@ -1 +0,0 @@
-b b/b
`,
			oldFilename: "b b/b",
			filename:    "b b/b",
			addition:    0,
			deletion:    1,
		},
		{
			name: "ambiguous addition",
			gitdiff: `diff --git a/b b/b b/b b/b
new file mode 100644
index 0000000..92e798b
--- /dev/null
+++ b/b b/b` + "\t" + `
@@ -0,0 +1 @@
+b b/b
`,
			oldFilename: "b b/b",
			filename:    "b b/b",
			addition:    1,
			deletion:    0,
		},
		{
			name: "rename",
			gitdiff: `diff --git a/b b/b b/b b/b b/b b/b
similarity index 100%
rename from b b/b b/b b/b b/b
rename to b
`,
			oldFilename: "b b/b b/b b/b b/b",
			filename:    "b",
		},
		{
			name: "ambiguous 1",
			gitdiff: `diff --git a/b b/b b/b b/b b/b b/b
similarity index 100%
rename from b b/b b/b b/b b/b
rename to b
`,
			oldFilename: "b b/b b/b b/b b/b",
			filename:    "b",
		},
		{
			name: "ambiguous 2",
			gitdiff: `diff --git a/b b/b b/b b/b b/b b/b
similarity index 100%
rename from b b/b b/b b/b
rename to b b/b
`,
			oldFilename: "b b/b b/b b/b",
			filename:    "b b/b",
		},
				t.Errorf("ParsePatch(%q) error = %v, wantErr %v", testcase.name, err, testcase.wantErr)
				t.Errorf("ParsePath(%q) did not receive 1 file:\n%s", testcase.name, string(gotMarshaled))
				t.Errorf("ParsePath(%q) does not have correct totalAddition %d, wanted %d", testcase.name, got.TotalAddition, testcase.addition)
				t.Errorf("ParsePath(%q) did not have correct totalDeletion %d, wanted %d", testcase.name, got.TotalDeletion, testcase.deletion)
				t.Errorf("ParsePath(%q) does not have correct file addition %d, wanted %d", testcase.name, file.Addition, testcase.addition)
				t.Errorf("ParsePath(%q) did not have correct file deletion %d, wanted %d", testcase.name, file.Deletion, testcase.deletion)
				t.Errorf("ParsePath(%q) did not have correct OldName %q, wanted %q", testcase.name, file.OldName, testcase.oldFilename)
				t.Errorf("ParsePath(%q) did not have correct Name %q, wanted %q", testcase.name, file.Name, testcase.filename)

func TestDiffToHTML_14231(t *testing.T) {
	setting.Cfg = ini.Empty()
	diffRecord := diffMatchPatch.DiffMain(highlight.Code("main.v", "		run()\n"), highlight.Code("main.v", "		run(db)\n"), true)
	diffRecord = diffMatchPatch.DiffCleanupEfficiency(diffRecord)

	expected := `		<span class="n">run</span><span class="added-code"><span class="o">(</span><span class="n">db</span></span><span class="o">)</span>`
	output := diffToHTML("main.v", diffRecord, DiffLineAdd)

	assertEqual(t, expected, output)
}