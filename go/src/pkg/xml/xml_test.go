<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/xml/xml_test.go</title>

  <link rel="stylesheet" type="text/css" href="../../../doc/style.css">
  <script type="text/javascript" src="../../../doc/godocs.js"></script>

</head>

<body>

  <script>
    // Catch 'enter' key down events and trigger the search form submission.
    function codesearchKeyDown(event) {
      if (event.which == 13) {
        var form = document.getElementById('codesearch');
        var query = document.getElementById('codesearchQuery');
        form.q.value = "lang:go package:go.googlecode.com " + query.value;
        document.getElementById('codesearch').submit();
}      return true;
}
    // Capture the submission event and construct the query parameter.
    function codeSearchSubmit() {
      var query = document.getElementById('codesearchQuery');
      var form = document.getElementById('codesearch');
      form.q.value = "lang:go package:go.googlecode.com " + query.value;
      return true;
}  </script>

<div id="topnav">
  <table summary="">
    <tr>
      <td id="headerImage">
        <a href="../../../index.html"><img src="../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
      </td>
      <td>
        <div id="headerDocSetTitle">The Go Programming Language</div>
      </td>
      <td>
        <!-- <table>
          <tr>
            <td>
              <! The input box is outside of the form because we want to add
              a couple of restricts to the query before submitting. If we just
              add the restricts to the text box before submitting, then they
              appear in the box when the user presses 'back'. Thus we use a
              hidden field in the form. However, there's no way to stop the
              non-hidden text box from also submitting a value unless we move
              it outside of the form
              <input type="search" id="codesearchQuery" value="" size="30" onkeydown="return codesearchKeyDown(event);"/>
              <form method="GET" action="http://www.google.com/codesearch" id="codesearch" class="search" onsubmit="return codeSearchSubmit();" style="display:inline;">
                <input type="hidden" name="q" value=""/>
                <input type="submit" value="Code search" />
                <span style="color: red">(TODO: remove for now?)</span>
              </form>
            </td>
          </tr>
          <tr>
            <td>
              <span style="color: gray;">(e.g. &ldquo;pem&rdquo; or &ldquo;xml&rdquo;)</span>
            </td>
          </tr>
        </table> -->
      </td>
    </tr>
  </table>
</div>

<div id="linkList">
  <ul>
    <li class="navhead"><a href="../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../index.html">Source files</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Help</li>
    <li>#go-nuts on irc.freenode.net</li>
    <li><a href="http://groups.google.com/group/golang-nuts">Go Nuts mailing list</a></li>
    <li><a href="http://code.google.com/p/go/issues/list">Issue tracker</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Go code search</li>
    <form method="GET" action="http://golang.org/search" class="search">
    <input type="search" name="q" value="" size="25" style="width:80%; max-width:200px" />
    <input type="submit" value="Go" />
    </form>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Last update</li>
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/xml/xml_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package xml

<a id="L7"></a>import (
    <a id="L8"></a>&#34;io&#34;;
    <a id="L9"></a>&#34;os&#34;;
    <a id="L10"></a>&#34;reflect&#34;;
    <a id="L11"></a>&#34;strings&#34;;
    <a id="L12"></a>&#34;testing&#34;;
<a id="L13"></a>)

<a id="L15"></a>const testInput = `
&lt;?xml version=&#34;1.0&#34; encoding=&#34;UTF-8&#34;?&gt;
&lt;!DOCTYPE html PUBLIC &#34;-//W3C//DTD XHTML 1.0 Transitional//EN&#34;
  &#34;http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd&#34;&gt;
&lt;body xmlns:foo=&#34;ns1&#34; xmlns=&#34;ns2&#34; xmlns:tag=&#34;ns3&#34; `
    <a id="L20"></a>&#34;\r\n\t&#34; `  &gt;
  &lt;hello lang=&#34;en&#34;&gt;World &amp;lt;&amp;gt;&amp;apos;&amp;quot; &amp;#x767d;&amp;#40300;翔&lt;/hello&gt;
  &lt;goodbye /&gt;
  &lt;outer foo:attr=&#34;value&#34; xmlns:tag=&#34;ns4&#34;&gt;
    &lt;inner/&gt;
  &lt;/outer&gt;
  &lt;tag:name&gt;
    Some text here.
  &lt;/tag:name&gt;
&lt;/body&gt;&lt;!-- missing final newline --&gt;`

<a id="L31"></a>var rawTokens = []Token{
    <a id="L32"></a>CharData(strings.Bytes(&#34;\n&#34;)),
    <a id="L33"></a>ProcInst{&#34;xml&#34;, strings.Bytes(`version=&#34;1.0&#34; encoding=&#34;UTF-8&#34;`)},
    <a id="L34"></a>CharData(strings.Bytes(&#34;\n&#34;)),
    <a id="L35"></a>Directive(strings.Bytes(`DOCTYPE html PUBLIC &#34;-//W3C//DTD XHTML 1.0 Transitional//EN&#34;
  &#34;http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd&#34;`<a id="L36"></a>)),
    <a id="L37"></a>CharData(strings.Bytes(&#34;\n&#34;)),
    <a id="L38"></a>StartElement{Name{&#34;&#34;, &#34;body&#34;}, []Attr{Attr{Name{&#34;xmlns&#34;, &#34;foo&#34;}, &#34;ns1&#34;}, Attr{Name{&#34;&#34;, &#34;xmlns&#34;}, &#34;ns2&#34;}, Attr{Name{&#34;xmlns&#34;, &#34;tag&#34;}, &#34;ns3&#34;}}},
    <a id="L39"></a>CharData(strings.Bytes(&#34;\n  &#34;)),
    <a id="L40"></a>StartElement{Name{&#34;&#34;, &#34;hello&#34;}, []Attr{Attr{Name{&#34;&#34;, &#34;lang&#34;}, &#34;en&#34;}}},
    <a id="L41"></a>CharData(strings.Bytes(&#34;World &lt;&gt;&#39;\&#34; 白鵬翔&#34;)),
    <a id="L42"></a>EndElement{Name{&#34;&#34;, &#34;hello&#34;}},
    <a id="L43"></a>CharData(strings.Bytes(&#34;\n  &#34;)),
    <a id="L44"></a>StartElement{Name{&#34;&#34;, &#34;goodbye&#34;}, nil},
    <a id="L45"></a>EndElement{Name{&#34;&#34;, &#34;goodbye&#34;}},
    <a id="L46"></a>CharData(strings.Bytes(&#34;\n  &#34;)),
    <a id="L47"></a>StartElement{Name{&#34;&#34;, &#34;outer&#34;}, []Attr{Attr{Name{&#34;foo&#34;, &#34;attr&#34;}, &#34;value&#34;}, Attr{Name{&#34;xmlns&#34;, &#34;tag&#34;}, &#34;ns4&#34;}}},
    <a id="L48"></a>CharData(strings.Bytes(&#34;\n    &#34;)),
    <a id="L49"></a>StartElement{Name{&#34;&#34;, &#34;inner&#34;}, nil},
    <a id="L50"></a>EndElement{Name{&#34;&#34;, &#34;inner&#34;}},
    <a id="L51"></a>CharData(strings.Bytes(&#34;\n  &#34;)),
    <a id="L52"></a>EndElement{Name{&#34;&#34;, &#34;outer&#34;}},
    <a id="L53"></a>CharData(strings.Bytes(&#34;\n  &#34;)),
    <a id="L54"></a>StartElement{Name{&#34;tag&#34;, &#34;name&#34;}, nil},
    <a id="L55"></a>CharData(strings.Bytes(&#34;\n    Some text here.\n  &#34;)),
    <a id="L56"></a>EndElement{Name{&#34;tag&#34;, &#34;name&#34;}},
    <a id="L57"></a>CharData(strings.Bytes(&#34;\n&#34;)),
    <a id="L58"></a>EndElement{Name{&#34;&#34;, &#34;body&#34;}},
    <a id="L59"></a>Comment(strings.Bytes(&#34; missing final newline &#34;)),
<a id="L60"></a>}

<a id="L62"></a>var cookedTokens = []Token{
    <a id="L63"></a>CharData(strings.Bytes(&#34;\n&#34;)),
    <a id="L64"></a>ProcInst{&#34;xml&#34;, strings.Bytes(`version=&#34;1.0&#34; encoding=&#34;UTF-8&#34;`)},
    <a id="L65"></a>CharData(strings.Bytes(&#34;\n&#34;)),
    <a id="L66"></a>Directive(strings.Bytes(`DOCTYPE html PUBLIC &#34;-//W3C//DTD XHTML 1.0 Transitional//EN&#34;
  &#34;http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd&#34;`<a id="L67"></a>)),
    <a id="L68"></a>CharData(strings.Bytes(&#34;\n&#34;)),
    <a id="L69"></a>StartElement{Name{&#34;ns2&#34;, &#34;body&#34;}, []Attr{Attr{Name{&#34;xmlns&#34;, &#34;foo&#34;}, &#34;ns1&#34;}, Attr{Name{&#34;&#34;, &#34;xmlns&#34;}, &#34;ns2&#34;}, Attr{Name{&#34;xmlns&#34;, &#34;tag&#34;}, &#34;ns3&#34;}}},
    <a id="L70"></a>CharData(strings.Bytes(&#34;\n  &#34;)),
    <a id="L71"></a>StartElement{Name{&#34;ns2&#34;, &#34;hello&#34;}, []Attr{Attr{Name{&#34;&#34;, &#34;lang&#34;}, &#34;en&#34;}}},
    <a id="L72"></a>CharData(strings.Bytes(&#34;World &lt;&gt;&#39;\&#34; 白鵬翔&#34;)),
    <a id="L73"></a>EndElement{Name{&#34;ns2&#34;, &#34;hello&#34;}},
    <a id="L74"></a>CharData(strings.Bytes(&#34;\n  &#34;)),
    <a id="L75"></a>StartElement{Name{&#34;ns2&#34;, &#34;goodbye&#34;}, nil},
    <a id="L76"></a>EndElement{Name{&#34;ns2&#34;, &#34;goodbye&#34;}},
    <a id="L77"></a>CharData(strings.Bytes(&#34;\n  &#34;)),
    <a id="L78"></a>StartElement{Name{&#34;ns2&#34;, &#34;outer&#34;}, []Attr{Attr{Name{&#34;ns1&#34;, &#34;attr&#34;}, &#34;value&#34;}, Attr{Name{&#34;xmlns&#34;, &#34;tag&#34;}, &#34;ns4&#34;}}},
    <a id="L79"></a>CharData(strings.Bytes(&#34;\n    &#34;)),
    <a id="L80"></a>StartElement{Name{&#34;ns2&#34;, &#34;inner&#34;}, nil},
    <a id="L81"></a>EndElement{Name{&#34;ns2&#34;, &#34;inner&#34;}},
    <a id="L82"></a>CharData(strings.Bytes(&#34;\n  &#34;)),
    <a id="L83"></a>EndElement{Name{&#34;ns2&#34;, &#34;outer&#34;}},
    <a id="L84"></a>CharData(strings.Bytes(&#34;\n  &#34;)),
    <a id="L85"></a>StartElement{Name{&#34;ns3&#34;, &#34;name&#34;}, nil},
    <a id="L86"></a>CharData(strings.Bytes(&#34;\n    Some text here.\n  &#34;)),
    <a id="L87"></a>EndElement{Name{&#34;ns3&#34;, &#34;name&#34;}},
    <a id="L88"></a>CharData(strings.Bytes(&#34;\n&#34;)),
    <a id="L89"></a>EndElement{Name{&#34;ns2&#34;, &#34;body&#34;}},
    <a id="L90"></a>Comment(strings.Bytes(&#34; missing final newline &#34;)),
<a id="L91"></a>}

<a id="L93"></a>type stringReader struct {
    <a id="L94"></a>s   string;
    <a id="L95"></a>off int;
<a id="L96"></a>}

<a id="L98"></a>func (r *stringReader) Read(b []byte) (n int, err os.Error) {
    <a id="L99"></a>if r.off &gt;= len(r.s) {
        <a id="L100"></a>return 0, os.EOF
    <a id="L101"></a>}
    <a id="L102"></a>for r.off &lt; len(r.s) &amp;&amp; n &lt; len(b) {
        <a id="L103"></a>b[n] = r.s[r.off];
        <a id="L104"></a>n++;
        <a id="L105"></a>r.off++;
    <a id="L106"></a>}
    <a id="L107"></a>return;
<a id="L108"></a>}

<a id="L110"></a>func (r *stringReader) ReadByte() (b byte, err os.Error) {
    <a id="L111"></a>if r.off &gt;= len(r.s) {
        <a id="L112"></a>return 0, os.EOF
    <a id="L113"></a>}
    <a id="L114"></a>b = r.s[r.off];
    <a id="L115"></a>r.off++;
    <a id="L116"></a>return;
<a id="L117"></a>}

<a id="L119"></a>func StringReader(s string) io.Reader { return &amp;stringReader{s, 0} }

<a id="L121"></a>func TestRawToken(t *testing.T) {
    <a id="L122"></a>p := NewParser(StringReader(testInput));

    <a id="L124"></a>for i, want := range rawTokens {
        <a id="L125"></a>have, err := p.RawToken();
        <a id="L126"></a>if err != nil {
            <a id="L127"></a>t.Fatalf(&#34;token %d: unexpected error: %s&#34;, i, err)
        <a id="L128"></a>}
        <a id="L129"></a>if !reflect.DeepEqual(have, want) {
            <a id="L130"></a>t.Errorf(&#34;token %d = %#v want %#v&#34;, i, have, want)
        <a id="L131"></a>}
    <a id="L132"></a>}
<a id="L133"></a>}

<a id="L135"></a>func TestToken(t *testing.T) {
    <a id="L136"></a>p := NewParser(StringReader(testInput));

    <a id="L138"></a>for i, want := range cookedTokens {
        <a id="L139"></a>have, err := p.Token();
        <a id="L140"></a>if err != nil {
            <a id="L141"></a>t.Fatalf(&#34;token %d: unexpected error: %s&#34;, i, err)
        <a id="L142"></a>}
        <a id="L143"></a>if !reflect.DeepEqual(have, want) {
            <a id="L144"></a>t.Errorf(&#34;token %d = %#v want %#v&#34;, i, have, want)
        <a id="L145"></a>}
    <a id="L146"></a>}
<a id="L147"></a>}
</pre>

</div>

<div id="footer">
<p>Except as noted, this content is
   licensed under <a href="http://creativecommons.org/licenses/by/3.0/">
   Creative Commons Attribution 3.0</a>.
</div>

<script type="text/javascript">
var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
var pageTracker = _gat._getTracker("UA-11222381-2");
pageTracker._trackPageview();
</script>
</body>
</html>
<!-- generated at Thu Nov 12 15:42:51 PST 2009 -->
