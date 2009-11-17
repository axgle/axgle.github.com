<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/godoc/main.go</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/cmd/godoc/main.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// godoc: Go Documentation Server</span>

<a id="L7"></a><span class="comment">// Web server tree:</span>
<a id="L8"></a><span class="comment">//</span>
<a id="L9"></a><span class="comment">//	http://godoc/		main landing page</span>
<a id="L10"></a><span class="comment">//	http://godoc/doc/	serve from $GOROOT/doc - spec, mem, tutorial, etc.</span>
<a id="L11"></a><span class="comment">//	http://godoc/src/	serve files from $GOROOT/src; .go gets pretty-printed</span>
<a id="L12"></a><span class="comment">//	http://godoc/cmd/	serve documentation about commands (TODO)</span>
<a id="L13"></a><span class="comment">//	http://godoc/pkg/	serve documentation about packages</span>
<a id="L14"></a><span class="comment">//				(idea is if you say import &#34;compress/zlib&#34;, you go to</span>
<a id="L15"></a><span class="comment">//				http://godoc/pkg/compress/zlib)</span>
<a id="L16"></a><span class="comment">//</span>
<a id="L17"></a><span class="comment">// Command-line interface:</span>
<a id="L18"></a><span class="comment">//</span>
<a id="L19"></a><span class="comment">//	godoc packagepath [name ...]</span>
<a id="L20"></a><span class="comment">//</span>
<a id="L21"></a><span class="comment">//	godoc compress/zlib</span>
<a id="L22"></a><span class="comment">//		- prints doc for package compress/zlib</span>
<a id="L23"></a><span class="comment">//	godoc crypto/block Cipher NewCMAC</span>
<a id="L24"></a><span class="comment">//		- prints doc for Cipher and NewCMAC in package crypto/block</span>

<a id="L26"></a>package main

<a id="L28"></a>import (
    <a id="L29"></a>&#34;bytes&#34;;
    <a id="L30"></a>&#34;flag&#34;;
    <a id="L31"></a>&#34;fmt&#34;;
    <a id="L32"></a>&#34;http&#34;;
    <a id="L33"></a>&#34;io&#34;;
    <a id="L34"></a>&#34;log&#34;;
    <a id="L35"></a>&#34;os&#34;;
    <a id="L36"></a>&#34;time&#34;;
<a id="L37"></a>)

<a id="L39"></a>var (
    <a id="L40"></a><span class="comment">// periodic sync</span>
    <a id="L41"></a>syncCmd              = flag.String(&#34;sync&#34;, &#34;&#34;, &#34;sync command; disabled if empty&#34;);
    <a id="L42"></a>syncMin              = flag.Int(&#34;sync_minutes&#34;, 0, &#34;sync interval in minutes; disabled if &lt;= 0&#34;);
    <a id="L43"></a>syncDelay delayTime; <span class="comment">// actual sync delay in minutes; usually syncDelay == syncMin, but delay may back off exponentially</span>

    <a id="L45"></a><span class="comment">// server control</span>
    <a id="L46"></a>httpaddr = flag.String(&#34;http&#34;, &#34;&#34;, &#34;HTTP service address (e.g., &#39;:6060&#39;)&#34;);

    <a id="L48"></a><span class="comment">// layout control</span>
    <a id="L49"></a>html = flag.Bool(&#34;html&#34;, false, &#34;print HTML in command-line mode&#34;);
<a id="L50"></a>)


<a id="L53"></a>func exec(c *http.Conn, args []string) (status int) {
    <a id="L54"></a>r, w, err := os.Pipe();
    <a id="L55"></a>if err != nil {
        <a id="L56"></a>log.Stderrf(&#34;os.Pipe(): %v\n&#34;, err);
        <a id="L57"></a>return 2;
    <a id="L58"></a>}

    <a id="L60"></a>bin := args[0];
    <a id="L61"></a>fds := []*os.File{nil, w, w};
    <a id="L62"></a>if *verbose {
        <a id="L63"></a>log.Stderrf(&#34;executing %v&#34;, args)
    <a id="L64"></a>}
    <a id="L65"></a>pid, err := os.ForkExec(bin, args, os.Environ(), goroot, fds);
    <a id="L66"></a>defer r.Close();
    <a id="L67"></a>w.Close();
    <a id="L68"></a>if err != nil {
        <a id="L69"></a>log.Stderrf(&#34;os.ForkExec(%q): %v\n&#34;, bin, err);
        <a id="L70"></a>return 2;
    <a id="L71"></a>}

    <a id="L73"></a>var buf bytes.Buffer;
    <a id="L74"></a>io.Copy(&amp;buf, r);
    <a id="L75"></a>wait, err := os.Wait(pid, 0);
    <a id="L76"></a>if err != nil {
        <a id="L77"></a>os.Stderr.Write(buf.Bytes());
        <a id="L78"></a>log.Stderrf(&#34;os.Wait(%d, 0): %v\n&#34;, pid, err);
        <a id="L79"></a>return 2;
    <a id="L80"></a>}
    <a id="L81"></a>status = wait.ExitStatus();
    <a id="L82"></a>if !wait.Exited() || status &gt; 1 {
        <a id="L83"></a>os.Stderr.Write(buf.Bytes());
        <a id="L84"></a>log.Stderrf(&#34;executing %v failed (exit status = %d)&#34;, args, status);
        <a id="L85"></a>return;
    <a id="L86"></a>}

    <a id="L88"></a>if *verbose {
        <a id="L89"></a>os.Stderr.Write(buf.Bytes())
    <a id="L90"></a>}
    <a id="L91"></a>if c != nil {
        <a id="L92"></a>c.SetHeader(&#34;content-type&#34;, &#34;text/plain; charset=utf-8&#34;);
        <a id="L93"></a>c.Write(buf.Bytes());
    <a id="L94"></a>}

    <a id="L96"></a>return;
<a id="L97"></a>}


<a id="L100"></a><span class="comment">// Maximum directory depth, adjust as needed.</span>
<a id="L101"></a>const maxDirDepth = 24

<a id="L103"></a>func dosync(c *http.Conn, r *http.Request) {
    <a id="L104"></a>args := []string{&#34;/bin/sh&#34;, &#34;-c&#34;, *syncCmd};
    <a id="L105"></a>switch exec(c, args) {
    <a id="L106"></a>case 0:
        <a id="L107"></a><span class="comment">// sync succeeded and some files have changed;</span>
        <a id="L108"></a><span class="comment">// update package tree.</span>
        <a id="L109"></a><span class="comment">// TODO(gri): The directory tree may be temporarily out-of-sync.</span>
        <a id="L110"></a><span class="comment">//            Consider keeping separate time stamps so the web-</span>
        <a id="L111"></a><span class="comment">//            page can indicate this discrepancy.</span>
        <a id="L112"></a>fsTree.set(newDirectory(&#34;.&#34;, maxDirDepth));
        <a id="L113"></a>fallthrough;
    <a id="L114"></a>case 1:
        <a id="L115"></a><span class="comment">// sync failed because no files changed;</span>
        <a id="L116"></a><span class="comment">// don&#39;t change the package tree</span>
        <a id="L117"></a>syncDelay.set(*syncMin) <span class="comment">//  revert to regular sync schedule</span>
    <a id="L118"></a>default:
        <a id="L119"></a><span class="comment">// sync failed because of an error - back off exponentially, but try at least once a day</span>
        <a id="L120"></a>syncDelay.backoff(24 * 60)
    <a id="L121"></a>}
<a id="L122"></a>}


<a id="L125"></a>func usage() {
    <a id="L126"></a>fmt.Fprintf(os.Stderr,
        <a id="L127"></a>&#34;usage: godoc package [name ...]\n&#34;
            <a id="L128"></a>&#34;	godoc -http=:6060\n&#34;);
    <a id="L129"></a>flag.PrintDefaults();
    <a id="L130"></a>os.Exit(2);
<a id="L131"></a>}


<a id="L134"></a>func loggingHandler(h http.Handler) http.Handler {
    <a id="L135"></a>return http.HandlerFunc(func(c *http.Conn, req *http.Request) {
        <a id="L136"></a>log.Stderrf(&#34;%s\t%s&#34;, c.RemoteAddr, req.URL);
        <a id="L137"></a>h.ServeHTTP(c, req);
    <a id="L138"></a>})
<a id="L139"></a>}


<a id="L142"></a>func main() {
    <a id="L143"></a>flag.Usage = usage;
    <a id="L144"></a>flag.Parse();

    <a id="L146"></a><span class="comment">// Check usage: either server and no args, or command line and args</span>
    <a id="L147"></a>if (*httpaddr != &#34;&#34;) != (flag.NArg() == 0) {
        <a id="L148"></a>usage()
    <a id="L149"></a>}

    <a id="L151"></a>if *tabwidth &lt; 0 {
        <a id="L152"></a>log.Exitf(&#34;negative tabwidth %d&#34;, *tabwidth)
    <a id="L153"></a>}

    <a id="L155"></a>if err := os.Chdir(goroot); err != nil {
        <a id="L156"></a>log.Exitf(&#34;chdir %s: %v&#34;, goroot, err)
    <a id="L157"></a>}

    <a id="L159"></a>readTemplates();

    <a id="L161"></a>if *httpaddr != &#34;&#34; {
        <a id="L162"></a><span class="comment">// HTTP server mode.</span>
        <a id="L163"></a>var handler http.Handler = http.DefaultServeMux;
        <a id="L164"></a>if *verbose {
            <a id="L165"></a>log.Stderrf(&#34;Go Documentation Server\n&#34;);
            <a id="L166"></a>log.Stderrf(&#34;address = %s\n&#34;, *httpaddr);
            <a id="L167"></a>log.Stderrf(&#34;goroot = %s\n&#34;, goroot);
            <a id="L168"></a>log.Stderrf(&#34;cmdroot = %s\n&#34;, *cmdroot);
            <a id="L169"></a>log.Stderrf(&#34;pkgroot = %s\n&#34;, *pkgroot);
            <a id="L170"></a>log.Stderrf(&#34;tmplroot = %s\n&#34;, *tmplroot);
            <a id="L171"></a>log.Stderrf(&#34;tabwidth = %d\n&#34;, *tabwidth);
            <a id="L172"></a>handler = loggingHandler(handler);
        <a id="L173"></a>}

        <a id="L175"></a>registerPublicHandlers(http.DefaultServeMux);
        <a id="L176"></a>if *syncCmd != &#34;&#34; {
            <a id="L177"></a>http.Handle(&#34;/debug/sync&#34;, http.HandlerFunc(dosync))
        <a id="L178"></a>}

        <a id="L180"></a><span class="comment">// Initialize directory tree with corresponding timestamp.</span>
        <a id="L181"></a><span class="comment">// Do it in two steps:</span>
        <a id="L182"></a><span class="comment">// 1) set timestamp right away so that the indexer is kicked on</span>
        <a id="L183"></a>fsTree.set(nil);
        <a id="L184"></a><span class="comment">// 2) compute initial directory tree in a goroutine so that launch is quick</span>
        <a id="L185"></a>go func() { fsTree.set(newDirectory(&#34;.&#34;, maxDirDepth)) }();

        <a id="L187"></a><span class="comment">// Start sync goroutine, if enabled.</span>
        <a id="L188"></a>if *syncCmd != &#34;&#34; &amp;&amp; *syncMin &gt; 0 {
            <a id="L189"></a>syncDelay.set(*syncMin); <span class="comment">// initial sync delay</span>
            <a id="L190"></a>go func() {
                <a id="L191"></a>for {
                    <a id="L192"></a>dosync(nil, nil);
                    <a id="L193"></a>delay, _ := syncDelay.get();
                    <a id="L194"></a>if *verbose {
                        <a id="L195"></a>log.Stderrf(&#34;next sync in %dmin&#34;, delay.(int))
                    <a id="L196"></a>}
                    <a id="L197"></a>time.Sleep(int64(delay.(int)) * 60e9);
                <a id="L198"></a>}
            <a id="L199"></a>}();
        <a id="L200"></a>}

        <a id="L202"></a><span class="comment">// Start indexing goroutine.</span>
        <a id="L203"></a>go indexer();

        <a id="L205"></a><span class="comment">// The server may have been restarted; always wait 1sec to</span>
        <a id="L206"></a><span class="comment">// give the forking server a chance to shut down and release</span>
        <a id="L207"></a><span class="comment">// the http port.</span>
        <a id="L208"></a><span class="comment">// TODO(gri): Do we still need this?</span>
        <a id="L209"></a>time.Sleep(1e9);

        <a id="L211"></a><span class="comment">// Start http server.</span>
        <a id="L212"></a>if err := http.ListenAndServe(*httpaddr, handler); err != nil {
            <a id="L213"></a>log.Exitf(&#34;ListenAndServe %s: %v&#34;, *httpaddr, err)
        <a id="L214"></a>}
        <a id="L215"></a>return;
    <a id="L216"></a>}

    <a id="L218"></a><span class="comment">// Command line mode.</span>
    <a id="L219"></a>if *html {
        <a id="L220"></a>packageText = packageHTML;
        <a id="L221"></a>parseerrorText = parseerrorHTML;
    <a id="L222"></a>}

    <a id="L224"></a>info := pkgHandler.getPageInfo(flag.Arg(0));

    <a id="L226"></a>if info.PDoc == nil &amp;&amp; info.Dirs == nil {
        <a id="L227"></a><span class="comment">// try again, this time assume it&#39;s a command</span>
        <a id="L228"></a>info = cmdHandler.getPageInfo(flag.Arg(0))
    <a id="L229"></a>}

    <a id="L231"></a>if info.PDoc != nil &amp;&amp; flag.NArg() &gt; 1 {
        <a id="L232"></a>args := flag.Args();
        <a id="L233"></a>info.PDoc.Filter(args[1:len(args)]);
    <a id="L234"></a>}

    <a id="L236"></a>if err := packageText.Execute(info, os.Stdout); err != nil {
        <a id="L237"></a>log.Stderrf(&#34;packageText.Execute: %s&#34;, err)
    <a id="L238"></a>}
<a id="L239"></a>}
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
