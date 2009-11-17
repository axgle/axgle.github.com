<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/http/triv.go</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/http/triv.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package main

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;bufio&#34;;
    <a id="L10"></a>&#34;expvar&#34;;
    <a id="L11"></a>&#34;flag&#34;;
    <a id="L12"></a>&#34;fmt&#34;;
    <a id="L13"></a>&#34;io&#34;;
    <a id="L14"></a>&#34;log&#34;;
    <a id="L15"></a>&#34;net&#34;;
    <a id="L16"></a>&#34;os&#34;;
    <a id="L17"></a>&#34;strconv&#34;;
<a id="L18"></a>)


<a id="L21"></a><span class="comment">// hello world, the web server</span>
<a id="L22"></a>var helloRequests = expvar.NewInt(&#34;hello-requests&#34;)

<a id="L24"></a>func HelloServer(c *http.Conn, req *http.Request) {
    <a id="L25"></a>helloRequests.Add(1);
    <a id="L26"></a>io.WriteString(c, &#34;hello, world!\n&#34;);
<a id="L27"></a>}

<a id="L29"></a><span class="comment">// Simple counter server. POSTing to it will set the value.</span>
<a id="L30"></a>type Counter struct {
    <a id="L31"></a>n int;
<a id="L32"></a>}

<a id="L34"></a><span class="comment">// This makes Counter satisfy the expvar.Var interface, so we can export</span>
<a id="L35"></a><span class="comment">// it directly.</span>
<a id="L36"></a>func (ctr *Counter) String() string { return fmt.Sprintf(&#34;%d&#34;, ctr.n) }

<a id="L38"></a>func (ctr *Counter) ServeHTTP(c *http.Conn, req *http.Request) {
    <a id="L39"></a>switch req.Method {
    <a id="L40"></a>case &#34;GET&#34;:
        <a id="L41"></a>ctr.n++
    <a id="L42"></a>case &#34;POST&#34;:
        <a id="L43"></a>buf := new(bytes.Buffer);
        <a id="L44"></a>io.Copy(buf, req.Body);
        <a id="L45"></a>body := buf.String();
        <a id="L46"></a>if n, err := strconv.Atoi(body); err != nil {
            <a id="L47"></a>fmt.Fprintf(c, &#34;bad POST: %v\nbody: [%v]\n&#34;, err, body)
        <a id="L48"></a>} else {
            <a id="L49"></a>ctr.n = n;
            <a id="L50"></a>fmt.Fprint(c, &#34;counter reset\n&#34;);
        <a id="L51"></a>}
    <a id="L52"></a>}
    <a id="L53"></a>fmt.Fprintf(c, &#34;counter = %d\n&#34;, ctr.n);
<a id="L54"></a>}

<a id="L56"></a><span class="comment">// simple file server</span>
<a id="L57"></a>var webroot = flag.String(&#34;root&#34;, &#34;/home/rsc&#34;, &#34;web root directory&#34;)
<a id="L58"></a>var pathVar = expvar.NewMap(&#34;file-requests&#34;)

<a id="L60"></a>func FileServer(c *http.Conn, req *http.Request) {
    <a id="L61"></a>c.SetHeader(&#34;content-type&#34;, &#34;text/plain; charset=utf-8&#34;);
    <a id="L62"></a>pathVar.Add(req.URL.Path, 1);
    <a id="L63"></a>path := *webroot + req.URL.Path; <span class="comment">// TODO: insecure: use os.CleanName</span>
    <a id="L64"></a>f, err := os.Open(path, os.O_RDONLY, 0);
    <a id="L65"></a>if err != nil {
        <a id="L66"></a>c.WriteHeader(http.StatusNotFound);
        <a id="L67"></a>fmt.Fprintf(c, &#34;open %s: %v\n&#34;, path, err);
        <a id="L68"></a>return;
    <a id="L69"></a>}
    <a id="L70"></a>n, err1 := io.Copy(c, f);
    <a id="L71"></a>fmt.Fprintf(c, &#34;[%d bytes]\n&#34;, n);
    <a id="L72"></a>f.Close();
<a id="L73"></a>}

<a id="L75"></a><span class="comment">// simple flag server</span>
<a id="L76"></a>var booleanflag = flag.Bool(&#34;boolean&#34;, true, &#34;another flag for testing&#34;)

<a id="L78"></a>func FlagServer(c *http.Conn, req *http.Request) {
    <a id="L79"></a>c.SetHeader(&#34;content-type&#34;, &#34;text/plain; charset=utf-8&#34;);
    <a id="L80"></a>fmt.Fprint(c, &#34;Flags:\n&#34;);
    <a id="L81"></a>flag.VisitAll(func(f *flag.Flag) {
        <a id="L82"></a>if f.Value.String() != f.DefValue {
            <a id="L83"></a>fmt.Fprintf(c, &#34;%s = %s [default = %s]\n&#34;, f.Name, f.Value.String(), f.DefValue)
        <a id="L84"></a>} else {
            <a id="L85"></a>fmt.Fprintf(c, &#34;%s = %s\n&#34;, f.Name, f.Value.String())
        <a id="L86"></a>}
    <a id="L87"></a>});
<a id="L88"></a>}

<a id="L90"></a><span class="comment">// simple argument server</span>
<a id="L91"></a>func ArgServer(c *http.Conn, req *http.Request) {
    <a id="L92"></a>for i, s := range os.Args {
        <a id="L93"></a>fmt.Fprint(c, s, &#34; &#34;)
    <a id="L94"></a>}
<a id="L95"></a>}

<a id="L97"></a><span class="comment">// a channel (just for the fun of it)</span>
<a id="L98"></a>type Chan chan int

<a id="L100"></a>func ChanCreate() Chan {
    <a id="L101"></a>c := make(Chan);
    <a id="L102"></a>go func(c Chan) {
        <a id="L103"></a>for x := 0; ; x++ {
            <a id="L104"></a>c &lt;- x
        <a id="L105"></a>}
    <a id="L106"></a>}(c);
    <a id="L107"></a>return c;
<a id="L108"></a>}

<a id="L110"></a>func (ch Chan) ServeHTTP(c *http.Conn, req *http.Request) {
    <a id="L111"></a>io.WriteString(c, fmt.Sprintf(&#34;channel send #%d\n&#34;, &lt;-ch))
<a id="L112"></a>}

<a id="L114"></a><span class="comment">// exec a program, redirecting output</span>
<a id="L115"></a>func DateServer(c *http.Conn, req *http.Request) {
    <a id="L116"></a>c.SetHeader(&#34;content-type&#34;, &#34;text/plain; charset=utf-8&#34;);
    <a id="L117"></a>r, w, err := os.Pipe();
    <a id="L118"></a>if err != nil {
        <a id="L119"></a>fmt.Fprintf(c, &#34;pipe: %s\n&#34;, err);
        <a id="L120"></a>return;
    <a id="L121"></a>}
    <a id="L122"></a>pid, err := os.ForkExec(&#34;/bin/date&#34;, []string{&#34;date&#34;}, os.Environ(), &#34;&#34;, []*os.File{nil, w, w});
    <a id="L123"></a>defer r.Close();
    <a id="L124"></a>w.Close();
    <a id="L125"></a>if err != nil {
        <a id="L126"></a>fmt.Fprintf(c, &#34;fork/exec: %s\n&#34;, err);
        <a id="L127"></a>return;
    <a id="L128"></a>}
    <a id="L129"></a>io.Copy(c, r);
    <a id="L130"></a>wait, err := os.Wait(pid, 0);
    <a id="L131"></a>if err != nil {
        <a id="L132"></a>fmt.Fprintf(c, &#34;wait: %s\n&#34;, err);
        <a id="L133"></a>return;
    <a id="L134"></a>}
    <a id="L135"></a>if !wait.Exited() || wait.ExitStatus() != 0 {
        <a id="L136"></a>fmt.Fprintf(c, &#34;date: %v\n&#34;, wait);
        <a id="L137"></a>return;
    <a id="L138"></a>}
<a id="L139"></a>}

<a id="L141"></a>func main() {
    <a id="L142"></a>flag.Parse();

    <a id="L144"></a><span class="comment">// The counter is published as a variable directly.</span>
    <a id="L145"></a>ctr := new(Counter);
    <a id="L146"></a>http.Handle(&#34;/counter&#34;, ctr);
    <a id="L147"></a>expvar.Publish(&#34;counter&#34;, ctr);

    <a id="L149"></a>http.Handle(&#34;/go/&#34;, http.HandlerFunc(FileServer));
    <a id="L150"></a>http.Handle(&#34;/flags&#34;, http.HandlerFunc(FlagServer));
    <a id="L151"></a>http.Handle(&#34;/args&#34;, http.HandlerFunc(ArgServer));
    <a id="L152"></a>http.Handle(&#34;/go/hello&#34;, http.HandlerFunc(HelloServer));
    <a id="L153"></a>http.Handle(&#34;/chan&#34;, ChanCreate());
    <a id="L154"></a>http.Handle(&#34;/date&#34;, http.HandlerFunc(DateServer));
    <a id="L155"></a>err := http.ListenAndServe(&#34;:12345&#34;, nil);
    <a id="L156"></a>if err != nil {
        <a id="L157"></a>log.Crash(&#34;ListenAndServe: &#34;, err)
    <a id="L158"></a>}
<a id="L159"></a>}
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
