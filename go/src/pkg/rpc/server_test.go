<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/rpc/server_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/rpc/server_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package rpc

<a id="L7"></a>import (
    <a id="L8"></a>&#34;http&#34;;
    <a id="L9"></a>&#34;log&#34;;
    <a id="L10"></a>&#34;net&#34;;
    <a id="L11"></a>&#34;once&#34;;
    <a id="L12"></a>&#34;os&#34;;
    <a id="L13"></a>&#34;strings&#34;;
    <a id="L14"></a>&#34;testing&#34;;
<a id="L15"></a>)

<a id="L17"></a>var serverAddr string
<a id="L18"></a>var httpServerAddr string

<a id="L20"></a>const second = 1e9


<a id="L23"></a>type Args struct {
    <a id="L24"></a>A, B int;
<a id="L25"></a>}

<a id="L27"></a>type Reply struct {
    <a id="L28"></a>C int;
<a id="L29"></a>}

<a id="L31"></a>type Arith int

<a id="L33"></a>func (t *Arith) Add(args *Args, reply *Reply) os.Error {
    <a id="L34"></a>reply.C = args.A + args.B;
    <a id="L35"></a>return nil;
<a id="L36"></a>}

<a id="L38"></a>func (t *Arith) Mul(args *Args, reply *Reply) os.Error {
    <a id="L39"></a>reply.C = args.A * args.B;
    <a id="L40"></a>return nil;
<a id="L41"></a>}

<a id="L43"></a>func (t *Arith) Div(args *Args, reply *Reply) os.Error {
    <a id="L44"></a>if args.B == 0 {
        <a id="L45"></a>return os.ErrorString(&#34;divide by zero&#34;)
    <a id="L46"></a>}
    <a id="L47"></a>reply.C = args.A / args.B;
    <a id="L48"></a>return nil;
<a id="L49"></a>}

<a id="L51"></a>func (t *Arith) Error(args *Args, reply *Reply) os.Error {
    <a id="L52"></a>panicln(&#34;ERROR&#34;)
<a id="L53"></a>}

<a id="L55"></a>func startServer() {
    <a id="L56"></a>Register(new(Arith));

    <a id="L58"></a>l, e := net.Listen(&#34;tcp&#34;, &#34;:0&#34;); <span class="comment">// any available address</span>
    <a id="L59"></a>if e != nil {
        <a id="L60"></a>log.Exitf(&#34;net.Listen tcp :0: %v&#34;, e)
    <a id="L61"></a>}
    <a id="L62"></a>serverAddr = l.Addr().String();
    <a id="L63"></a>log.Stderr(&#34;Test RPC server listening on &#34;, serverAddr);
    <a id="L64"></a>go Accept(l);

    <a id="L66"></a>HandleHTTP();
    <a id="L67"></a>l, e = net.Listen(&#34;tcp&#34;, &#34;:0&#34;); <span class="comment">// any available address</span>
    <a id="L68"></a>if e != nil {
        <a id="L69"></a>log.Stderrf(&#34;net.Listen tcp :0: %v&#34;, e);
        <a id="L70"></a>os.Exit(1);
    <a id="L71"></a>}
    <a id="L72"></a>httpServerAddr = l.Addr().String();
    <a id="L73"></a>log.Stderr(&#34;Test HTTP RPC server listening on &#34;, httpServerAddr);
    <a id="L74"></a>go http.Serve(l, nil);
<a id="L75"></a>}

<a id="L77"></a>func TestRPC(t *testing.T) {
    <a id="L78"></a>once.Do(startServer);

    <a id="L80"></a>client, err := Dial(&#34;tcp&#34;, serverAddr);
    <a id="L81"></a>if err != nil {
        <a id="L82"></a>t.Fatal(&#34;dialing&#34;, err)
    <a id="L83"></a>}

    <a id="L85"></a><span class="comment">// Synchronous calls</span>
    <a id="L86"></a>args := &amp;Args{7, 8};
    <a id="L87"></a>reply := new(Reply);
    <a id="L88"></a>err = client.Call(&#34;Arith.Add&#34;, args, reply);
    <a id="L89"></a>if reply.C != args.A+args.B {
        <a id="L90"></a>t.Errorf(&#34;Add: expected %d got %d&#34;, reply.C, args.A+args.B)
    <a id="L91"></a>}

    <a id="L93"></a>args = &amp;Args{7, 8};
    <a id="L94"></a>reply = new(Reply);
    <a id="L95"></a>err = client.Call(&#34;Arith.Mul&#34;, args, reply);
    <a id="L96"></a>if reply.C != args.A*args.B {
        <a id="L97"></a>t.Errorf(&#34;Mul: expected %d got %d&#34;, reply.C, args.A*args.B)
    <a id="L98"></a>}

    <a id="L100"></a><span class="comment">// Out of order.</span>
    <a id="L101"></a>args = &amp;Args{7, 8};
    <a id="L102"></a>mulReply := new(Reply);
    <a id="L103"></a>mulCall := client.Go(&#34;Arith.Mul&#34;, args, mulReply, nil);
    <a id="L104"></a>addReply := new(Reply);
    <a id="L105"></a>addCall := client.Go(&#34;Arith.Add&#34;, args, addReply, nil);

    <a id="L107"></a>&lt;-addCall.Done;
    <a id="L108"></a>if addReply.C != args.A+args.B {
        <a id="L109"></a>t.Errorf(&#34;Add: expected %d got %d&#34;, addReply.C, args.A+args.B)
    <a id="L110"></a>}

    <a id="L112"></a>&lt;-mulCall.Done;
    <a id="L113"></a>if mulReply.C != args.A*args.B {
        <a id="L114"></a>t.Errorf(&#34;Mul: expected %d got %d&#34;, mulReply.C, args.A*args.B)
    <a id="L115"></a>}

    <a id="L117"></a><span class="comment">// Error test</span>
    <a id="L118"></a>args = &amp;Args{7, 0};
    <a id="L119"></a>reply = new(Reply);
    <a id="L120"></a>err = client.Call(&#34;Arith.Div&#34;, args, reply);
    <a id="L121"></a><span class="comment">// expect an error: zero divide</span>
    <a id="L122"></a>if err == nil {
        <a id="L123"></a>t.Error(&#34;Div: expected error&#34;)
    <a id="L124"></a>} else if err.String() != &#34;divide by zero&#34; {
        <a id="L125"></a>t.Error(&#34;Div: expected divide by zero error; got&#34;, err)
    <a id="L126"></a>}
<a id="L127"></a>}

<a id="L129"></a>func TestHTTPRPC(t *testing.T) {
    <a id="L130"></a>once.Do(startServer);

    <a id="L132"></a>client, err := DialHTTP(&#34;tcp&#34;, httpServerAddr);
    <a id="L133"></a>if err != nil {
        <a id="L134"></a>t.Fatal(&#34;dialing&#34;, err)
    <a id="L135"></a>}

    <a id="L137"></a><span class="comment">// Synchronous calls</span>
    <a id="L138"></a>args := &amp;Args{7, 8};
    <a id="L139"></a>reply := new(Reply);
    <a id="L140"></a>err = client.Call(&#34;Arith.Add&#34;, args, reply);
    <a id="L141"></a>if reply.C != args.A+args.B {
        <a id="L142"></a>t.Errorf(&#34;Add: expected %d got %d&#34;, reply.C, args.A+args.B)
    <a id="L143"></a>}
<a id="L144"></a>}

<a id="L146"></a>func TestCheckUnknownService(t *testing.T) {
    <a id="L147"></a>once.Do(startServer);

    <a id="L149"></a>conn, err := net.Dial(&#34;tcp&#34;, &#34;&#34;, serverAddr);
    <a id="L150"></a>if err != nil {
        <a id="L151"></a>t.Fatal(&#34;dialing:&#34;, err)
    <a id="L152"></a>}

    <a id="L154"></a>client := NewClient(conn);

    <a id="L156"></a>args := &amp;Args{7, 8};
    <a id="L157"></a>reply := new(Reply);
    <a id="L158"></a>err = client.Call(&#34;Unknown.Add&#34;, args, reply);
    <a id="L159"></a>if err == nil {
        <a id="L160"></a>t.Error(&#34;expected error calling unknown service&#34;)
    <a id="L161"></a>} else if strings.Index(err.String(), &#34;service&#34;) &lt; 0 {
        <a id="L162"></a>t.Error(&#34;expected error about service; got&#34;, err)
    <a id="L163"></a>}
<a id="L164"></a>}

<a id="L166"></a>func TestCheckUnknownMethod(t *testing.T) {
    <a id="L167"></a>once.Do(startServer);

    <a id="L169"></a>conn, err := net.Dial(&#34;tcp&#34;, &#34;&#34;, serverAddr);
    <a id="L170"></a>if err != nil {
        <a id="L171"></a>t.Fatal(&#34;dialing:&#34;, err)
    <a id="L172"></a>}

    <a id="L174"></a>client := NewClient(conn);

    <a id="L176"></a>args := &amp;Args{7, 8};
    <a id="L177"></a>reply := new(Reply);
    <a id="L178"></a>err = client.Call(&#34;Arith.Unknown&#34;, args, reply);
    <a id="L179"></a>if err == nil {
        <a id="L180"></a>t.Error(&#34;expected error calling unknown service&#34;)
    <a id="L181"></a>} else if strings.Index(err.String(), &#34;method&#34;) &lt; 0 {
        <a id="L182"></a>t.Error(&#34;expected error about method; got&#34;, err)
    <a id="L183"></a>}
<a id="L184"></a>}

<a id="L186"></a>func TestCheckBadType(t *testing.T) {
    <a id="L187"></a>once.Do(startServer);

    <a id="L189"></a>conn, err := net.Dial(&#34;tcp&#34;, &#34;&#34;, serverAddr);
    <a id="L190"></a>if err != nil {
        <a id="L191"></a>t.Fatal(&#34;dialing:&#34;, err)
    <a id="L192"></a>}

    <a id="L194"></a>client := NewClient(conn);

    <a id="L196"></a>reply := new(Reply);
    <a id="L197"></a>err = client.Call(&#34;Arith.Add&#34;, reply, reply); <span class="comment">// args, reply would be the correct thing to use</span>
    <a id="L198"></a>if err == nil {
        <a id="L199"></a>t.Error(&#34;expected error calling Arith.Add with wrong arg type&#34;)
    <a id="L200"></a>} else if strings.Index(err.String(), &#34;type&#34;) &lt; 0 {
        <a id="L201"></a>t.Error(&#34;expected error about type; got&#34;, err)
    <a id="L202"></a>}
<a id="L203"></a>}
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
