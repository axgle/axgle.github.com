<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/proc/regs_linux_amd64.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../doc/style.css">
  <script type="text/javascript" src="../../../../doc/godocs.js"></script>

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
        <a href="../../../../index.html"><img src="../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/debug/proc/regs_linux_amd64.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package proc

<a id="L7"></a>import (
    <a id="L8"></a>&#34;os&#34;;
    <a id="L9"></a>&#34;strconv&#34;;
    <a id="L10"></a>&#34;syscall&#34;;
<a id="L11"></a>)

<a id="L13"></a>type amd64Regs struct {
    <a id="L14"></a>syscall.PtraceRegs;
    <a id="L15"></a>setter func(*syscall.PtraceRegs) os.Error;
<a id="L16"></a>}

<a id="L18"></a>var names = [...]string{
    <a id="L19"></a>&#34;rax&#34;,
    <a id="L20"></a>&#34;rbx&#34;,
    <a id="L21"></a>&#34;rcx&#34;,
    <a id="L22"></a>&#34;rdx&#34;,
    <a id="L23"></a>&#34;rsi&#34;,
    <a id="L24"></a>&#34;rdi&#34;,
    <a id="L25"></a>&#34;rbp&#34;,
    <a id="L26"></a>&#34;rsp&#34;,
    <a id="L27"></a>&#34;r8&#34;,
    <a id="L28"></a>&#34;r9&#34;,
    <a id="L29"></a>&#34;r10&#34;,
    <a id="L30"></a>&#34;r11&#34;,
    <a id="L31"></a>&#34;r12&#34;,
    <a id="L32"></a>&#34;r13&#34;,
    <a id="L33"></a>&#34;r14&#34;,
    <a id="L34"></a>&#34;r15&#34;,
    <a id="L35"></a>&#34;rip&#34;,
    <a id="L36"></a>&#34;eflags&#34;,
    <a id="L37"></a>&#34;cs&#34;,
    <a id="L38"></a>&#34;ss&#34;,
    <a id="L39"></a>&#34;ds&#34;,
    <a id="L40"></a>&#34;es&#34;,
    <a id="L41"></a>&#34;fs&#34;,
    <a id="L42"></a>&#34;gs&#34;,

    <a id="L44"></a><span class="comment">// PtraceRegs contains these registers, but I don&#39;t think</span>
    <a id="L45"></a><span class="comment">// they&#39;re actually meaningful.</span>
    <a id="L46"></a><span class="comment">//&#34;orig_rax&#34;,</span>
    <a id="L47"></a><span class="comment">//&#34;fs_base&#34;,</span>
    <a id="L48"></a><span class="comment">//&#34;gs_base&#34;,</span>
<a id="L49"></a>}

<a id="L51"></a>func (r *amd64Regs) PC() Word { return Word(r.Rip) }

<a id="L53"></a>func (r *amd64Regs) SetPC(val Word) os.Error {
    <a id="L54"></a>r.Rip = uint64(val);
    <a id="L55"></a>return r.setter(&amp;r.PtraceRegs);
<a id="L56"></a>}

<a id="L58"></a>func (r *amd64Regs) Link() Word {
    <a id="L59"></a><span class="comment">// TODO(austin)</span>
    <a id="L60"></a>panic(&#34;No link register&#34;)
<a id="L61"></a>}

<a id="L63"></a>func (r *amd64Regs) SetLink(val Word) os.Error {
    <a id="L64"></a>panic(&#34;No link register&#34;)
<a id="L65"></a>}

<a id="L67"></a>func (r *amd64Regs) SP() Word { return Word(r.Rsp) }

<a id="L69"></a>func (r *amd64Regs) SetSP(val Word) os.Error {
    <a id="L70"></a>r.Rsp = uint64(val);
    <a id="L71"></a>return r.setter(&amp;r.PtraceRegs);
<a id="L72"></a>}

<a id="L74"></a>func (r *amd64Regs) Names() []string { return &amp;names }

<a id="L76"></a>func (r *amd64Regs) Get(i int) Word {
    <a id="L77"></a>switch i {
    <a id="L78"></a>case 0:
        <a id="L79"></a>return Word(r.Rax)
    <a id="L80"></a>case 1:
        <a id="L81"></a>return Word(r.Rbx)
    <a id="L82"></a>case 2:
        <a id="L83"></a>return Word(r.Rcx)
    <a id="L84"></a>case 3:
        <a id="L85"></a>return Word(r.Rdx)
    <a id="L86"></a>case 4:
        <a id="L87"></a>return Word(r.Rsi)
    <a id="L88"></a>case 5:
        <a id="L89"></a>return Word(r.Rdi)
    <a id="L90"></a>case 6:
        <a id="L91"></a>return Word(r.Rbp)
    <a id="L92"></a>case 7:
        <a id="L93"></a>return Word(r.Rsp)
    <a id="L94"></a>case 8:
        <a id="L95"></a>return Word(r.R8)
    <a id="L96"></a>case 9:
        <a id="L97"></a>return Word(r.R9)
    <a id="L98"></a>case 10:
        <a id="L99"></a>return Word(r.R10)
    <a id="L100"></a>case 11:
        <a id="L101"></a>return Word(r.R11)
    <a id="L102"></a>case 12:
        <a id="L103"></a>return Word(r.R12)
    <a id="L104"></a>case 13:
        <a id="L105"></a>return Word(r.R13)
    <a id="L106"></a>case 14:
        <a id="L107"></a>return Word(r.R14)
    <a id="L108"></a>case 15:
        <a id="L109"></a>return Word(r.R15)
    <a id="L110"></a>case 16:
        <a id="L111"></a>return Word(r.Rip)
    <a id="L112"></a>case 17:
        <a id="L113"></a>return Word(r.Eflags)
    <a id="L114"></a>case 18:
        <a id="L115"></a>return Word(r.Cs)
    <a id="L116"></a>case 19:
        <a id="L117"></a>return Word(r.Ss)
    <a id="L118"></a>case 20:
        <a id="L119"></a>return Word(r.Ds)
    <a id="L120"></a>case 21:
        <a id="L121"></a>return Word(r.Es)
    <a id="L122"></a>case 22:
        <a id="L123"></a>return Word(r.Fs)
    <a id="L124"></a>case 23:
        <a id="L125"></a>return Word(r.Gs)
    <a id="L126"></a>}
    <a id="L127"></a>panic(&#34;invalid register index &#34;, strconv.Itoa(i));
<a id="L128"></a>}

<a id="L130"></a>func (r *amd64Regs) Set(i int, val Word) os.Error {
    <a id="L131"></a>switch i {
    <a id="L132"></a>case 0:
        <a id="L133"></a>r.Rax = uint64(val)
    <a id="L134"></a>case 1:
        <a id="L135"></a>r.Rbx = uint64(val)
    <a id="L136"></a>case 2:
        <a id="L137"></a>r.Rcx = uint64(val)
    <a id="L138"></a>case 3:
        <a id="L139"></a>r.Rdx = uint64(val)
    <a id="L140"></a>case 4:
        <a id="L141"></a>r.Rsi = uint64(val)
    <a id="L142"></a>case 5:
        <a id="L143"></a>r.Rdi = uint64(val)
    <a id="L144"></a>case 6:
        <a id="L145"></a>r.Rbp = uint64(val)
    <a id="L146"></a>case 7:
        <a id="L147"></a>r.Rsp = uint64(val)
    <a id="L148"></a>case 8:
        <a id="L149"></a>r.R8 = uint64(val)
    <a id="L150"></a>case 9:
        <a id="L151"></a>r.R9 = uint64(val)
    <a id="L152"></a>case 10:
        <a id="L153"></a>r.R10 = uint64(val)
    <a id="L154"></a>case 11:
        <a id="L155"></a>r.R11 = uint64(val)
    <a id="L156"></a>case 12:
        <a id="L157"></a>r.R12 = uint64(val)
    <a id="L158"></a>case 13:
        <a id="L159"></a>r.R13 = uint64(val)
    <a id="L160"></a>case 14:
        <a id="L161"></a>r.R14 = uint64(val)
    <a id="L162"></a>case 15:
        <a id="L163"></a>r.R15 = uint64(val)
    <a id="L164"></a>case 16:
        <a id="L165"></a>r.Rip = uint64(val)
    <a id="L166"></a>case 17:
        <a id="L167"></a>r.Eflags = uint64(val)
    <a id="L168"></a>case 18:
        <a id="L169"></a>r.Cs = uint64(val)
    <a id="L170"></a>case 19:
        <a id="L171"></a>r.Ss = uint64(val)
    <a id="L172"></a>case 20:
        <a id="L173"></a>r.Ds = uint64(val)
    <a id="L174"></a>case 21:
        <a id="L175"></a>r.Es = uint64(val)
    <a id="L176"></a>case 22:
        <a id="L177"></a>r.Fs = uint64(val)
    <a id="L178"></a>case 23:
        <a id="L179"></a>r.Gs = uint64(val)
    <a id="L180"></a>default:
        <a id="L181"></a>panic(&#34;invalid register index &#34;, strconv.Itoa(i))
    <a id="L182"></a>}
    <a id="L183"></a>return r.setter(&amp;r.PtraceRegs);
<a id="L184"></a>}

<a id="L186"></a>func newRegs(regs *syscall.PtraceRegs, setter func(*syscall.PtraceRegs) os.Error) Regs {
    <a id="L187"></a>res := amd64Regs{};
    <a id="L188"></a>res.PtraceRegs = *regs;
    <a id="L189"></a>res.setter = setter;
    <a id="L190"></a>return &amp;res;
<a id="L191"></a>}
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
