支持渲染 Markdown

例如：

```bash
$ tree
./example/
├── markdown.get.json
├── README.md
$ cat ./example/markdown.get.json
{
  "body": "markdown://./README.md"
}
$ gmock ./
$ curl http://localhost:8080/markdown
<h3 id="hello-you-are-access">Hello, you are access</h3>

<p>this is a markdown render</p>

<ul>
<li>[x] You you like it?</li>
<li>[x] yes.</li>
</ul>

<table>
<thead>
<tr>
<th>Name</th>
<th>Age</th>
</tr>
</thead>

<tbody>
<tr>
<td>Bob</td>
<td>27</td>
</tr>

<tr>
<td>Alice</td>
<td>23</td>
</tr>
</tbody>
</table>

<pre><code>func getTrue() bool {
    return true
}
</code></pre>

<dl>
<dt>Cat</dt>
<dd>Fluffy animal everyone likes</dd>
<dt>Internet</dt>
<dd>Vector of transmission for pictures of cats</dd>
</dl>

<p>This is a footnote.[^1]</p>

<p>[^1]: the footnote text.</p>
```