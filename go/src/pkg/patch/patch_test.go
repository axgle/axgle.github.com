<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/patch/patch_test.go</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/patch/patch_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package patch

<a id="L7"></a><span class="comment">// TODO(rsc): test Apply</span>

<a id="L9"></a>import (
    <a id="L10"></a>&#34;strings&#34;;
    <a id="L11"></a>&#34;testing&#34;;
<a id="L12"></a>)

<a id="L14"></a>type Test struct {
    <a id="L15"></a>in   string;
    <a id="L16"></a>out  string;
    <a id="L17"></a>diff string;
<a id="L18"></a>}

<a id="L20"></a>func TestFileApply(t *testing.T) {
    <a id="L21"></a>for i, test := range tests {
        <a id="L22"></a>set, err := Parse(strings.Bytes(test.diff));
        <a id="L23"></a>if err != nil {
            <a id="L24"></a>t.Errorf(&#34;#%d: Parse: %s&#34;, i, err);
            <a id="L25"></a>continue;
        <a id="L26"></a>}
        <a id="L27"></a>if len(set.File) != 1 {
            <a id="L28"></a>t.Errorf(&#34;#%d: Parse returned %d patches, want 1&#34;, i, len(set.File));
            <a id="L29"></a>continue;
        <a id="L30"></a>}
        <a id="L31"></a>new, err := set.File[0].Apply(strings.Bytes(test.in));
        <a id="L32"></a>if err != nil {
            <a id="L33"></a>t.Errorf(&#34;#%d: Apply: %s&#34;, i, err);
            <a id="L34"></a>continue;
        <a id="L35"></a>}
        <a id="L36"></a>if s := string(new); s != test.out {
            <a id="L37"></a>t.Errorf(&#34;#%d:\n--- have\n%s--- want\n%s&#34;, i, s, test.out)
        <a id="L38"></a>}
    <a id="L39"></a>}
<a id="L40"></a>}

<a id="L42"></a>var tests = []Test{
    <a id="L43"></a>Test{
        <a id="L44"></a>&#34;hello, world\n&#34;,
        <a id="L45"></a>&#34;goodbye, world\n&#34;,
        <a id="L46"></a>&#34;Index: a\n&#34;
            <a id="L47"></a>&#34;--- a/a\n&#34;
            <a id="L48"></a>&#34;+++ b/b\n&#34;
            <a id="L49"></a>&#34;@@ -1 +1 @@\n&#34;
            <a id="L50"></a>&#34;-hello, world\n&#34;
            <a id="L51"></a>&#34;+goodbye, world\n&#34;,
    <a id="L52"></a>},
    <a id="L53"></a>Test{
        <a id="L54"></a>&#34;hello, world\n&#34;,
        <a id="L55"></a>&#34;goodbye, world\n&#34;,
        <a id="L56"></a>&#34;diff a/a b/b\n&#34;
            <a id="L57"></a>&#34;--- a/a\n&#34;
            <a id="L58"></a>&#34;+++ b/b\n&#34;
            <a id="L59"></a>&#34;@@ -1,1 +1,1 @@\n&#34;
            <a id="L60"></a>&#34;-hello, world\n&#34;
            <a id="L61"></a>&#34;+goodbye, world\n&#34;,
    <a id="L62"></a>},
    <a id="L63"></a>Test{
        <a id="L64"></a>&#34;hello, world&#34;,
        <a id="L65"></a>&#34;goodbye, world\n&#34;,
        <a id="L66"></a>&#34;diff --git a/a b/b\n&#34;
            <a id="L67"></a>&#34;--- a/a\n&#34;
            <a id="L68"></a>&#34;+++ b/b\n&#34;
            <a id="L69"></a>&#34;@@ -1 +1 @@\n&#34;
            <a id="L70"></a>&#34;-hello, world\n&#34;
            <a id="L71"></a>&#34;\\ No newline at end of file\n&#34;
            <a id="L72"></a>&#34;+goodbye, world\n&#34;,
    <a id="L73"></a>},
    <a id="L74"></a>Test{
        <a id="L75"></a>&#34;hello, world\n&#34;,
        <a id="L76"></a>&#34;goodbye, world&#34;,
        <a id="L77"></a>&#34;Index: a\n&#34;
            <a id="L78"></a>&#34;--- a/a\n&#34;
            <a id="L79"></a>&#34;+++ b/b\n&#34;
            <a id="L80"></a>&#34;@@ -1 +1 @@\n&#34;
            <a id="L81"></a>&#34;-hello, world\n&#34;
            <a id="L82"></a>&#34;+goodbye, world\n&#34;
            <a id="L83"></a>&#34;\\ No newline at end of file\n&#34;,
    <a id="L84"></a>},
    <a id="L85"></a>Test{
        <a id="L86"></a>&#34;hello, world&#34;,
        <a id="L87"></a>&#34;goodbye, world&#34;,
        <a id="L88"></a>&#34;Index: a\n&#34;
            <a id="L89"></a>&#34;--- a/a\n&#34;
            <a id="L90"></a>&#34;+++ b/b\n&#34;
            <a id="L91"></a>&#34;@@ -1 +1 @@\n&#34;
            <a id="L92"></a>&#34;-hello, world\n&#34;
            <a id="L93"></a>&#34;\\ No newline at end of file\n&#34;
            <a id="L94"></a>&#34;+goodbye, world\n&#34;
            <a id="L95"></a>&#34;\\ No newline at end of file\n&#34;,
    <a id="L96"></a>},
    <a id="L97"></a>Test{
        <a id="L98"></a>&#34;a\nb\nc\nd\ne\nf\ng\nh\ni\nj\nk\nl\nm\nn\n&#34;,
        <a id="L99"></a>&#34;a\nB\nC\nD\ne\nf\ng\nj\nk\nl\nm\nN\n&#34;,
        <a id="L100"></a>&#34;Index: a\n&#34;
            <a id="L101"></a>&#34;--- a/a\n&#34;
            <a id="L102"></a>&#34;+++ b/b\n&#34;
            <a id="L103"></a>&#34;@@ -1,14 +1,12 @@\n&#34;
            <a id="L104"></a>&#34; a\n&#34;
            <a id="L105"></a>&#34;-b\n&#34;
            <a id="L106"></a>&#34;-c\n&#34;
            <a id="L107"></a>&#34;-d\n&#34;
            <a id="L108"></a>&#34;+B\n&#34;
            <a id="L109"></a>&#34;+C\n&#34;
            <a id="L110"></a>&#34;+D\n&#34;
            <a id="L111"></a>&#34; e\n&#34;
            <a id="L112"></a>&#34; f\n&#34;
            <a id="L113"></a>&#34; g\n&#34;
            <a id="L114"></a>&#34;-h\n&#34;
            <a id="L115"></a>&#34;-i\n&#34;
            <a id="L116"></a>&#34; j\n&#34;
            <a id="L117"></a>&#34; k\n&#34;
            <a id="L118"></a>&#34; l\n&#34;
            <a id="L119"></a>&#34; m\n&#34;
            <a id="L120"></a>&#34;-n\n&#34;
            <a id="L121"></a>&#34;+N\n&#34;,
    <a id="L122"></a>},
    <a id="L123"></a>Test{
        <a id="L124"></a>&#34;a\nb\nc\nd\ne\nf\ng\nh\ni\nj\nk\nl\nm\nn\no\np\nq\nr\ns\nt\nu\nv\nw\nx\ny\nz\n&#34;,
        <a id="L125"></a>&#34;a\nb\nc\ng\nh\ni\nj\nk\nl\nm\nN\nO\np\nq\nr\ns\nt\nu\nv\nw\nd\ne\nf\nx\n&#34;,
        <a id="L126"></a>&#34;Index: a\n&#34;
            <a id="L127"></a>&#34;--- a/a\n&#34;
            <a id="L128"></a>&#34;+++ b/b\n&#34;
            <a id="L129"></a>&#34;@@ -1,9 +1,6 @@\n&#34;
            <a id="L130"></a>&#34; a\n&#34;
            <a id="L131"></a>&#34; b\n&#34;
            <a id="L132"></a>&#34; c\n&#34;
            <a id="L133"></a>&#34;-d\n&#34;
            <a id="L134"></a>&#34;-e\n&#34;
            <a id="L135"></a>&#34;-f\n&#34;
            <a id="L136"></a>&#34; g\n&#34;
            <a id="L137"></a>&#34; h\n&#34;
            <a id="L138"></a>&#34; i\n&#34;
            <a id="L139"></a>&#34;@@ -11,8 +8,8 @@ j\n&#34;
            <a id="L140"></a>&#34; k\n&#34;
            <a id="L141"></a>&#34; l\n&#34;
            <a id="L142"></a>&#34; m\n&#34;
            <a id="L143"></a>&#34;-n\n&#34;
            <a id="L144"></a>&#34;-o\n&#34;
            <a id="L145"></a>&#34;+N\n&#34;
            <a id="L146"></a>&#34;+O\n&#34;
            <a id="L147"></a>&#34; p\n&#34;
            <a id="L148"></a>&#34; q\n&#34;
            <a id="L149"></a>&#34; r\n&#34;
            <a id="L150"></a>&#34;\n&#34;
            <a id="L151"></a>&#34;@@ -21,6 +18,7 @@ t\n&#34;
            <a id="L152"></a>&#34; u\n&#34;
            <a id="L153"></a>&#34; v\n&#34;
            <a id="L154"></a>&#34; w\n&#34;
            <a id="L155"></a>&#34;+d\n&#34;
            <a id="L156"></a>&#34;+e\n&#34;
            <a id="L157"></a>&#34;+f\n&#34;
            <a id="L158"></a>&#34; x\n&#34;
            <a id="L159"></a>&#34;-y\n&#34;
            <a id="L160"></a>&#34;-z\n&#34;,
    <a id="L161"></a>},
    <a id="L162"></a>Test{
        <a id="L163"></a>&#34;a\nb\nc\ng\nh\ni\nj\nk\nl\nm\nN\nO\np\nq\nr\ns\nt\nu\nv\nw\nd\ne\nf\nx\n&#34;,
        <a id="L164"></a>&#34;a\nb\nc\nd\ne\nf\ng\nh\ni\nj\nk\nl\nm\nn\no\np\nq\nr\ns\nt\nu\nv\nw\nx\ny\nz\n&#34;,
        <a id="L165"></a>&#34;Index: a\n&#34;
            <a id="L166"></a>&#34;--- a/b\n&#34;
            <a id="L167"></a>&#34;+++ b/a\n&#34;
            <a id="L168"></a>&#34;@@ -1,6 +1,9 @@\n&#34;
            <a id="L169"></a>&#34; a\n&#34;
            <a id="L170"></a>&#34; b\n&#34;
            <a id="L171"></a>&#34; c\n&#34;
            <a id="L172"></a>&#34;+d\n&#34;
            <a id="L173"></a>&#34;+e\n&#34;
            <a id="L174"></a>&#34;+f\n&#34;
            <a id="L175"></a>&#34; g\n&#34;
            <a id="L176"></a>&#34; h\n&#34;
            <a id="L177"></a>&#34; i\n&#34;
            <a id="L178"></a>&#34;@@ -8,8 +11,8 @@ j\n&#34;
            <a id="L179"></a>&#34; k\n&#34;
            <a id="L180"></a>&#34; l\n&#34;
            <a id="L181"></a>&#34; m\n&#34;
            <a id="L182"></a>&#34;-N\n&#34;
            <a id="L183"></a>&#34;-O\n&#34;
            <a id="L184"></a>&#34;+n\n&#34;
            <a id="L185"></a>&#34;+o\n&#34;
            <a id="L186"></a>&#34; p\n&#34;
            <a id="L187"></a>&#34; q\n&#34;
            <a id="L188"></a>&#34; r\n&#34;
            <a id="L189"></a>&#34;@@ -18,7 +21,6 @@ t\n&#34;
            <a id="L190"></a>&#34; u\n&#34;
            <a id="L191"></a>&#34; v\n&#34;
            <a id="L192"></a>&#34; w\n&#34;
            <a id="L193"></a>&#34;-d\n&#34;
            <a id="L194"></a>&#34;-e\n&#34;
            <a id="L195"></a>&#34;-f\n&#34;
            <a id="L196"></a>&#34; x\n&#34;
            <a id="L197"></a>&#34;+y\n&#34;
            <a id="L198"></a>&#34;+z\n&#34;,
    <a id="L199"></a>},
    <a id="L200"></a>Test{
        <a id="L201"></a>&#34;a\nb\nc\nd\ne\nf\ng\nh\ni\nj\nk\nl\nm\nn\no\np\nq\nr\ns\nt\nu\nv\nw\nx\ny\nz\n&#34;,
        <a id="L202"></a>&#34;&#34;,
        <a id="L203"></a>&#34;Index: a\n&#34;
            <a id="L204"></a>&#34;deleted file mode 100644\n&#34;
            <a id="L205"></a>&#34;--- a/a\n&#34;
            <a id="L206"></a>&#34;+++ /dev/null\n&#34;
            <a id="L207"></a>&#34;@@ -1,26 +0,0 @@\n&#34;
            <a id="L208"></a>&#34;-a\n&#34;
            <a id="L209"></a>&#34;-b\n&#34;
            <a id="L210"></a>&#34;-c\n&#34;
            <a id="L211"></a>&#34;-d\n&#34;
            <a id="L212"></a>&#34;-e\n&#34;
            <a id="L213"></a>&#34;-f\n&#34;
            <a id="L214"></a>&#34;-g\n&#34;
            <a id="L215"></a>&#34;-h\n&#34;
            <a id="L216"></a>&#34;-i\n&#34;
            <a id="L217"></a>&#34;-j\n&#34;
            <a id="L218"></a>&#34;-k\n&#34;
            <a id="L219"></a>&#34;-l\n&#34;
            <a id="L220"></a>&#34;-m\n&#34;
            <a id="L221"></a>&#34;-n\n&#34;
            <a id="L222"></a>&#34;-o\n&#34;
            <a id="L223"></a>&#34;-p\n&#34;
            <a id="L224"></a>&#34;-q\n&#34;
            <a id="L225"></a>&#34;-r\n&#34;
            <a id="L226"></a>&#34;-s\n&#34;
            <a id="L227"></a>&#34;-t\n&#34;
            <a id="L228"></a>&#34;-u\n&#34;
            <a id="L229"></a>&#34;-v\n&#34;
            <a id="L230"></a>&#34;-w\n&#34;
            <a id="L231"></a>&#34;-x\n&#34;
            <a id="L232"></a>&#34;-y\n&#34;
            <a id="L233"></a>&#34;-z\n&#34;,
    <a id="L234"></a>},
    <a id="L235"></a>Test{
        <a id="L236"></a>&#34;&#34;,
        <a id="L237"></a>&#34;a\nb\nc\nd\ne\nf\ng\nh\ni\nj\nk\nl\nm\nn\no\np\nq\nr\ns\nt\nu\nv\nw\nx\ny\nz\n&#34;,
        <a id="L238"></a>&#34;Index: a\n&#34;
            <a id="L239"></a>&#34;new file mode 100644\n&#34;
            <a id="L240"></a>&#34;--- /dev/null\n&#34;
            <a id="L241"></a>&#34;+++ b/a\n&#34;
            <a id="L242"></a>&#34;@@ -0,0 +1,26 @@\n&#34;
            <a id="L243"></a>&#34;+a\n&#34;
            <a id="L244"></a>&#34;+b\n&#34;
            <a id="L245"></a>&#34;+c\n&#34;
            <a id="L246"></a>&#34;+d\n&#34;
            <a id="L247"></a>&#34;+e\n&#34;
            <a id="L248"></a>&#34;+f\n&#34;
            <a id="L249"></a>&#34;+g\n&#34;
            <a id="L250"></a>&#34;+h\n&#34;
            <a id="L251"></a>&#34;+i\n&#34;
            <a id="L252"></a>&#34;+j\n&#34;
            <a id="L253"></a>&#34;+k\n&#34;
            <a id="L254"></a>&#34;+l\n&#34;
            <a id="L255"></a>&#34;+m\n&#34;
            <a id="L256"></a>&#34;+n\n&#34;
            <a id="L257"></a>&#34;+o\n&#34;
            <a id="L258"></a>&#34;+p\n&#34;
            <a id="L259"></a>&#34;+q\n&#34;
            <a id="L260"></a>&#34;+r\n&#34;
            <a id="L261"></a>&#34;+s\n&#34;
            <a id="L262"></a>&#34;+t\n&#34;
            <a id="L263"></a>&#34;+u\n&#34;
            <a id="L264"></a>&#34;+v\n&#34;
            <a id="L265"></a>&#34;+w\n&#34;
            <a id="L266"></a>&#34;+x\n&#34;
            <a id="L267"></a>&#34;+y\n&#34;
            <a id="L268"></a>&#34;+z\n&#34;,
    <a id="L269"></a>},
    <a id="L270"></a>Test{
        <a id="L271"></a>&#34;\xc2\xd8\xf9\x63\x8c\xf7\xc6\x9b\xb0\x3c\x39\xfa\x08\x8e\x42\x8f&#34;
            <a id="L272"></a>&#34;\x1c\x7c\xaf\x54\x22\x87\xc3\xc5\x68\x9b\xe1\xbd\xbc\xc3\xe0\xda&#34;
            <a id="L273"></a>&#34;\xcc\xe3\x96\xda\xc2\xaf\xbb\x75\x79\x64\x86\x60\x8a\x43\x9e\x07&#34;
            <a id="L274"></a>&#34;\x9c\xaa\x92\x88\xd4\x30\xb9\x8b\x95\x04\x60\x71\xc7\xbb\x2d\x93&#34;
            <a id="L275"></a>&#34;\x66\x73\x01\x24\xf3\x63\xbf\xe6\x1d\x38\x15\x56\x98\xc4\x1f\x85&#34;
            <a id="L276"></a>&#34;\xc3\x60\x39\x3a\x0d\x57\x53\x0c\x29\x3f\xbb\x44\x7e\x56\x56\x9d&#34;
            <a id="L277"></a>&#34;\x87\xcf\xf6\x88\xe8\x98\x05\x85\xf8\xfe\x44\x21\xfa\x33\xc9\xa4&#34;
            <a id="L278"></a>&#34;\x22\xbe\x89\x05\x8b\x82\x76\xc9\x7c\xaf\x48\x28\xc4\x86\x15\x89&#34;
            <a id="L279"></a>&#34;\xb9\x98\xfa\x41\xfc\x3d\x8d\x80\x29\x33\x17\x45\xa5\x7f\x67\x79&#34;
            <a id="L280"></a>&#34;\x7f\x92\x3b\x2e\x4c\xc1\xd2\x1b\x9e\xcf\xed\x53\x56\xb2\x49\x58&#34;
            <a id="L281"></a>&#34;\xd8\xe9\x9f\x98\xa3\xfe\x78\xe1\xe8\x74\x71\x04\x1a\x87\xd9\x68&#34;
            <a id="L282"></a>&#34;\x18\x68\xd0\xae\x7b\xa4\x25\xe3\x06\x03\x7e\x8b\xd3\x50\x1f\xb1&#34;
            <a id="L283"></a>&#34;\x67\x08\xe3\x93\xf4\x4f\xa1\xfb\x31\xcf\x99\x5a\x43\x9f\x4b\xc4&#34;
            <a id="L284"></a>&#34;\xaa\x68\x1a\xf9\x8e\x97\x02\x80\x17\xf1\x25\x21\xdf\x94\xbf\x41&#34;
            <a id="L285"></a>&#34;\x08\x59\x3d\xea\x36\x23\x03\xb5\x62\x4d\xb6\x8f\x9e\xdf\x1f\x03&#34;
            <a id="L286"></a>&#34;\x7d\x70\xe0\x6f\x46\x08\x96\x79\x72\xb7\xae\x41\x2b\xbd\x2a\x95&#34;,

        <a id="L288"></a>&#34;\x8e\x5f\xf8\x79\x36\x8d\xbe\x68\xc4\x2c\x78\x8a\x46\x28\x40\x3e&#34;
            <a id="L289"></a>&#34;\xcf\x3b\xb9\x14\xaf\xfa\x04\x9e\x4b\xa2\x52\x51\x51\xf0\xad\xd3&#34;
            <a id="L290"></a>&#34;\x03\x1c\x03\x79\x5f\x53\xc7\x1a\xd5\x28\xe2\xd9\x19\x37\xa4\xfa&#34;
            <a id="L291"></a>&#34;\xdd\xff\xac\xb5\xa9\x42\x4e\x17\xeb\xb4\x0d\x20\x67\x08\x43\x21&#34;
            <a id="L292"></a>&#34;\x7d\x12\x27\xfa\x96\x7a\x85\xf8\x04\x5f\xf4\xfe\xda\x9f\x66\xf2&#34;
            <a id="L293"></a>&#34;\xba\x04\x39\x00\xab\x3f\x23\x20\x84\x53\xb4\x88\xb6\xee\xa2\x9e&#34;
            <a id="L294"></a>&#34;\xc1\xca\xd4\x09\x2a\x27\x89\x2f\xcb\xba\xa6\x41\xb6\xe9\xc5\x08&#34;
            <a id="L295"></a>&#34;\xff\xf5\x95\x35\xab\xbb\x5c\x62\x96\xe7\x7c\x8f\xf2\x40\x12\xc9&#34;
            <a id="L296"></a>&#34;\x2d\xfe\xff\x75\x4f\x70\x47\xc9\xcd\x15\x0a\x1c\x23\xe7\x0f\x15&#34;
            <a id="L297"></a>&#34;\x95\x75\x30\x8f\x6e\x9f\x7e\xa5\x9d\xd1\x65\x1c\x4d\x4e\xf4\x32&#34;
            <a id="L298"></a>&#34;\x49\x9b\xa1\x30\x44\x62\x6f\xe2\xe6\x69\x09\xf8\x7c\x7c\xbe\x07&#34;
            <a id="L299"></a>&#34;\xa9\xb6\x14\x7a\x6b\x85\xe4\xbf\x48\xbe\x5b\x3b\x70\xb3\x79\x3b&#34;
            <a id="L300"></a>&#34;\xc4\x35\x9d\x86\xf1\xfe\x2b\x6f\x80\x74\x50\xf3\x96\x59\x53\x1a&#34;
            <a id="L301"></a>&#34;\x75\x46\x9d\x57\x72\xb3\xb1\x26\xf5\x81\xcd\x96\x08\xbc\x2b\x10&#34;
            <a id="L302"></a>&#34;\xdc\x80\xbd\xd0\xdf\x03\x6d\x8d\xec\x30\x2b\x4c\xdb\x4d\x3b\xef&#34;
            <a id="L303"></a>&#34;\x7d\x3a\x39\xc8\x5a\xc4\xcc\x24\x37\xde\xe2\x95\x2b\x04\x97\xb0&#34;,

        <a id="L305"></a><span class="comment">// From git diff --binary</span>
        <a id="L306"></a>&#34;Index: a\n&#34;
            <a id="L307"></a>&#34;index cb34d9b1743b7c410fa750be8a58eb355987110b..0a01764bc1b2fd29da317f72208f462ad342400f 100644\n&#34;
            <a id="L308"></a>&#34;GIT binary patch\n&#34;
            <a id="L309"></a>&#34;literal 256\n&#34;
            <a id="L310"></a>&#34;zcmV+b0ssDvU-)@8jlO8aEO?4WC_p~XJGm6E`UIX!qEb;&amp;@U7DW90Pe@Q^y+BDB{@}\n&#34;
            <a id="L311"></a>&#34;zH&gt;CRA|E#sCLQWU!v&lt;)C&lt;2ty%#5-0kWdWHA|U-bUkpJwv91UUe!KO-Q7Q?!V-?xLQ-\n&#34;
            <a id="L312"></a>&#34;z%G3!eCy6i1x~4(4&gt;BR{D^_4ZNyIf+H=X{UyKoZF&lt;{{MAPa7W3_6$%_9=MNQ?buf=^\n&#34;
            <a id="L313"></a>&#34;zpMIsC(PbP&gt;PV_QKo1rj7VsGN+X$kmze7*;%wiJ46h2+0TzFRwRvw1tjHJyg&gt;{wr^Q\n&#34;
            <a id="L314"></a>&#34;zbWrn_SyLKyMx9r3v#}=ifz6f(yekmgfW6S)18t4$Fe^;kO*`*&gt;IyuN%#LOf&amp;-r|)j\n&#34;
            <a id="L315"></a>&#34;G1edVN^?m&amp;S\n&#34;
            <a id="L316"></a>&#34;\n&#34;
            <a id="L317"></a>&#34;literal 256\n&#34;
            <a id="L318"></a>&#34;zcmV+b0ssEO*!g3O_r{yBJURLZjzW(de6Lg@hr`8ao8i5@!{FM?&lt;CfaOue)`5WQJgh\n&#34;
            <a id="L319"></a>&#34;zL!Jkms*;G*Fu9AB1YmK;yDgJua{(mtW54DdI2Bfy#2&lt;yjU^zMsS5pirKf6SJR#u&amp;d\n&#34;
            <a id="L320"></a>&#34;z&amp;-RGum&lt;5IS{zM`AGs&amp;bPzKI2kf_BM#uSh7wh82mqnEFBdJ&amp;k}VGZ#gre`k4rk~=O;\n&#34;
            <a id="L321"></a>&#34;z!O|O^&amp;+SuIvPoFj&gt;7SUR{&amp;?Z&amp;ba4b4huLTtXwa^Eq$T491AdFsP#&gt;{p2;-CVPoeuU\n&#34;
            <a id="L322"></a>&#34;z&amp;zV|7pG(B5Xd3yBmjZwn@g*VOl)pg;Sv~4DBLlT!O}3Ao-yZ{gaNuu72$p$rx2{1e\n&#34;
            <a id="L323"></a>&#34;Gy(*Pb;D3Ms\n&#34;
            <a id="L324"></a>&#34;\n&#34;,
    <a id="L325"></a>},
    <a id="L326"></a>Test{
        <a id="L327"></a>&#34;\xc2\xd8\xf9\x63\x8c\xf7\xc6\x9b\xb0\x3c\x39\xfa\x08\x8e\x42\x8f&#34;
            <a id="L328"></a>&#34;\x1c\x7c\xaf\x54\x22\x87\xc3\xc5\x68\x9b\xe1\xbd\xbc\xc3\xe0\xda&#34;
            <a id="L329"></a>&#34;\xcc\xe3\x96\xda\xc2\xaf\xbb\x75\x79\x64\x86\x60\x8a\x43\x9e\x07&#34;
            <a id="L330"></a>&#34;\x9c\xaa\x92\x88\xd4\x30\xb9\x8b\x95\x04\x60\x71\xc7\xbb\x2d\x93&#34;
            <a id="L331"></a>&#34;\x66\x73\x01\x24\xf3\x63\xbf\xe6\x1d\x38\x15\x56\x98\xc4\x1f\x85&#34;
            <a id="L332"></a>&#34;\xc3\x60\x39\x3a\x0d\x57\x53\x0c\x29\x3f\xbb\x44\x7e\x56\x56\x9d&#34;
            <a id="L333"></a>&#34;\x87\xcf\xf6\x88\xe8\x98\x05\x85\xf8\xfe\x44\x21\xfa\x33\xc9\xa4&#34;
            <a id="L334"></a>&#34;\x22\xbe\x89\x05\x8b\x82\x76\xc9\x7c\xaf\x48\x28\xc4\x86\x15\x89&#34;
            <a id="L335"></a>&#34;\xb9\x98\xfa\x41\xfc\x3d\x8d\x80\x29\x33\x17\x45\xa5\x7f\x67\x79&#34;
            <a id="L336"></a>&#34;\x7f\x92\x3b\x2e\x4c\xc1\xd2\x1b\x9e\xcf\xed\x53\x56\xb2\x49\x58&#34;
            <a id="L337"></a>&#34;\xd8\xe9\x9f\x98\xa3\xfe\x78\xe1\xe8\x74\x71\x04\x1a\x87\xd9\x68&#34;
            <a id="L338"></a>&#34;\x18\x68\xd0\xae\x7b\xa4\x25\xe3\x06\x03\x7e\x8b\xd3\x50\x1f\xb1&#34;
            <a id="L339"></a>&#34;\x67\x08\xe3\x93\xf4\x4f\xa1\xfb\x31\xcf\x99\x5a\x43\x9f\x4b\xc4&#34;
            <a id="L340"></a>&#34;\xaa\x68\x1a\xf9\x8e\x97\x02\x80\x17\xf1\x25\x21\xdf\x94\xbf\x41&#34;
            <a id="L341"></a>&#34;\x08\x59\x3d\xea\x36\x23\x03\xb5\x62\x4d\xb6\x8f\x9e\xdf\x1f\x03&#34;
            <a id="L342"></a>&#34;\x7d\x70\xe0\x6f\x46\x08\x96\x79\x72\xb7\xae\x41\x2b\xbd\x2a\x95&#34;,

        <a id="L344"></a>&#34;\x8e\x5f\xf8\x79\x36\x8d\xbe\x68\xc4\x2c\x78\x8a\x46\x28\x40\x3e&#34;
            <a id="L345"></a>&#34;\xcf\x3b\xb9\x14\xaf\xfa\x04\x9e\x4b\xa2\x52\x51\x51\xf0\xad\xd3&#34;
            <a id="L346"></a>&#34;\x03\x1c\x03\x79\x5f\x53\xc7\x1a\xd5\x28\xe2\xd9\x19\x37\xa4\xfa&#34;
            <a id="L347"></a>&#34;\xdd\xff\xac\xb5\xa9\x42\x4e\x17\xeb\xb4\x0d\x20\x67\x08\x43\x21&#34;
            <a id="L348"></a>&#34;\x7d\x12\x27\xfa\x96\x7a\x85\xf8\x04\x5f\xf4\xfe\xda\x9f\x66\xf2&#34;
            <a id="L349"></a>&#34;\xba\x04\x39\x00\xab\x3f\x23\x20\x84\x53\xb4\x88\xb6\xee\xa2\x9e&#34;
            <a id="L350"></a>&#34;\xc1\xca\xd4\x09\x2a\x27\x89\x2f\xcb\xba\xa6\x41\xb6\xe9\xc5\x08&#34;
            <a id="L351"></a>&#34;\xff\xf5\x95\x35\xab\xbb\x5c\x62\x96\xe7\x7c\x8f\xf2\x40\x12\xc9&#34;
            <a id="L352"></a>&#34;\x2d\xfe\xff\x75\x4f\x70\x47\xc9\xcd\x15\x0a\x1c\x23\xe7\x0f\x15&#34;
            <a id="L353"></a>&#34;\x95\x75\x30\x8f\x6e\x9f\x7e\xa5\x9d\xd1\x65\x1c\x4d\x4e\xf4\x32&#34;
            <a id="L354"></a>&#34;\x49\x9b\xa1\x30\x44\x62\x6f\xe2\xe6\x69\x09\xf8\x7c\x7c\xbe\x07&#34;
            <a id="L355"></a>&#34;\xa9\xb6\x14\x7a\x6b\x85\xe4\xbf\x48\xbe\x5b\x3b\x70\xb3\x79\x3b&#34;
            <a id="L356"></a>&#34;\xc4\x35\x9d\x86\xf1\xfe\x2b\x6f\x80\x74\x50\xf3\x96\x59\x53\x1a&#34;
            <a id="L357"></a>&#34;\x75\x46\x9d\x57\x72\xb3\xb1\x26\xf5\x81\xcd\x96\x08\xbc\x2b\x10&#34;
            <a id="L358"></a>&#34;\xdc\x80\xbd\xd0\xdf\x03\x6d\x8d\xec\x30\x2b\x4c\xdb\x4d\x3b\xef&#34;
            <a id="L359"></a>&#34;\x7d\x3a\x39\xc8\x5a\xc4\xcc\x24\x37\xde\xe2\x95\x2b\x04\x97\xb0&#34;,

        <a id="L361"></a><span class="comment">// From hg diff --git</span>
        <a id="L362"></a>&#34;Index: a\n&#34;
            <a id="L363"></a>&#34;index cb34d9b1743b7c410fa750be8a58eb355987110b..0a01764bc1b2fd29da317f72208f462ad342400f\n&#34;
            <a id="L364"></a>&#34;GIT binary patch\n&#34;
            <a id="L365"></a>&#34;literal 256\n&#34;
            <a id="L366"></a>&#34;zc$@(M0ssDvU-)@8jlO8aEO?4WC_p~XJGm6E`UIX!qEb;&amp;@U7DW90Pe@Q^y+BDB{@}\n&#34;
            <a id="L367"></a>&#34;zH&gt;CRA|E#sCLQWU!v&lt;)C&lt;2ty%#5-0kWdWHA|U-bUkpJwv91UUe!KO-Q7Q?!V-?xLQ-\n&#34;
            <a id="L368"></a>&#34;z%G3!eCy6i1x~4(4&gt;BR{D^_4ZNyIf+H=X{UyKoZF&lt;{{MAPa7W3_6$%_9=MNQ?buf=^\n&#34;
            <a id="L369"></a>&#34;zpMIsC(PbP&gt;PV_QKo1rj7VsGN+X$kmze7*;%wiJ46h2+0TzFRwRvw1tjHJyg&gt;{wr^Q\n&#34;
            <a id="L370"></a>&#34;zbWrn_SyLKyMx9r3v#}=ifz6f(yekmgfW6S)18t4$Fe^;kO*`*&gt;IyuN%#LOf&amp;-r|)j\n&#34;
            <a id="L371"></a>&#34;G1edVN^?m&amp;S\n&#34;
            <a id="L372"></a>&#34;\n&#34;,
    <a id="L373"></a>},
    <a id="L374"></a>Test{
        <a id="L375"></a>&#34;&#34;,
        <a id="L376"></a>&#34;&#34;,
        <a id="L377"></a>&#34;Index: hello\n&#34;
            <a id="L378"></a>&#34;===================================================================\n&#34;
            <a id="L379"></a>&#34;old mode 100644\n&#34;
            <a id="L380"></a>&#34;new mode 100755\n&#34;,
    <a id="L381"></a>},
<a id="L382"></a>}
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
