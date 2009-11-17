<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/fmt/print.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/fmt/print.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">/*</span>
<a id="L6"></a><span class="comment">	Package fmt implements formatted I/O with functions analogous</span>
<a id="L7"></a><span class="comment">	to C&#39;s printf.  The format &#39;verbs&#39; are derived from C&#39;s but</span>
<a id="L8"></a><span class="comment">	are simpler.</span>

<a id="L10"></a><span class="comment">	The verbs:</span>

<a id="L12"></a><span class="comment">	General:</span>
<a id="L13"></a><span class="comment">		%v	the value in a default format.</span>
<a id="L14"></a><span class="comment">			when printing structs, the plus flag (%+v) adds field names</span>
<a id="L15"></a><span class="comment">		%#v	a Go-syntax representation of the value</span>
<a id="L16"></a><span class="comment">		%T	a Go-syntax representation of the type of the value</span>

<a id="L18"></a><span class="comment">	Boolean:</span>
<a id="L19"></a><span class="comment">		%t	the word true or false</span>
<a id="L20"></a><span class="comment">	Integer:</span>
<a id="L21"></a><span class="comment">		%b	base 2</span>
<a id="L22"></a><span class="comment">		%c	the character represented by the corresponding Unicode code point</span>
<a id="L23"></a><span class="comment">		%d	base 10</span>
<a id="L24"></a><span class="comment">		%o	base 8</span>
<a id="L25"></a><span class="comment">		%x	base 16, with lower-case letters for a-f</span>
<a id="L26"></a><span class="comment">		%X	base 16, with upper-case letters for A-F</span>
<a id="L27"></a><span class="comment">	Floating-point:</span>
<a id="L28"></a><span class="comment">		%e	scientific notation, e.g. -1234.456e+78</span>
<a id="L29"></a><span class="comment">		%E	scientific notation, e.g. -1234.456E+78</span>
<a id="L30"></a><span class="comment">		%f	decimal point but no exponent, e.g. 123.456</span>
<a id="L31"></a><span class="comment">		%g	whichever of %e or %f produces more compact output</span>
<a id="L32"></a><span class="comment">		%G	whichever of %E or %f produces more compact output</span>
<a id="L33"></a><span class="comment">	String and slice of bytes:</span>
<a id="L34"></a><span class="comment">		%s	the uninterpreted bytes of the string or slice</span>
<a id="L35"></a><span class="comment">		%q	a double-quoted string safely escaped with Go syntax</span>
<a id="L36"></a><span class="comment">		%x	base 16 notation with two characters per byte</span>
<a id="L37"></a><span class="comment">	Pointer:</span>
<a id="L38"></a><span class="comment">		%p	base 16 notation, with leading 0x</span>

<a id="L40"></a><span class="comment">	There is no &#39;u&#39; flag.  Integers are printed unsigned if they have unsigned type.</span>
<a id="L41"></a><span class="comment">	Similarly, there is no need to specify the size of the operand (int8, int64).</span>

<a id="L43"></a><span class="comment">	For numeric values, the width and precision flags control</span>
<a id="L44"></a><span class="comment">	formatting; width sets the width of the field, precision the</span>
<a id="L45"></a><span class="comment">	number of places after the decimal, if appropriate.  The</span>
<a id="L46"></a><span class="comment">	format %6.2f prints 123.45.</span>

<a id="L48"></a><span class="comment">	Other flags:</span>
<a id="L49"></a><span class="comment">		+	always print a sign for numeric values</span>
<a id="L50"></a><span class="comment">		-	pad with spaces on the right rather than the left (left-justify the field)</span>
<a id="L51"></a><span class="comment">		#	alternate format: add leading 0 for octal (%#o), 0x for hex (%#x);</span>
<a id="L52"></a><span class="comment">			suppress 0x for %p (%#p);</span>
<a id="L53"></a><span class="comment">			print a raw (backquoted) string if possible for %q (%#q)</span>
<a id="L54"></a><span class="comment">		&#39; &#39;	(space) leave a space for elided sign in numbers (% d);</span>
<a id="L55"></a><span class="comment">			put spaces between bytes printing strings or slices in hex (% x)</span>
<a id="L56"></a><span class="comment">		0	pad with leading zeros rather than spaces</span>

<a id="L58"></a><span class="comment">	For each Printf-like function, there is also a Print function</span>
<a id="L59"></a><span class="comment">	that takes no format and is equivalent to saying %v for every</span>
<a id="L60"></a><span class="comment">	operand.  Another variant Println inserts blanks between</span>
<a id="L61"></a><span class="comment">	operands and appends a newline.</span>

<a id="L63"></a><span class="comment">	Regardless of the verb, if an operand is an interface value,</span>
<a id="L64"></a><span class="comment">	the internal concrete value is used, not the interface itself.</span>
<a id="L65"></a><span class="comment">	Thus:</span>
<a id="L66"></a><span class="comment">		var i interface{} = 23;</span>
<a id="L67"></a><span class="comment">		fmt.Printf(&#34;%v\n&#34;, i);</span>
<a id="L68"></a><span class="comment">	will print 23.</span>

<a id="L70"></a><span class="comment">	If an operand implements interface Formatter, that interface</span>
<a id="L71"></a><span class="comment">	can be used for fine control of formatting.</span>

<a id="L73"></a><span class="comment">	If an operand implements method String() string that method</span>
<a id="L74"></a><span class="comment">	will be used for %v, %s, or Print etc.</span>
<a id="L75"></a><span class="comment">*/</span>
<a id="L76"></a>package fmt


<a id="L79"></a>import (
    <a id="L80"></a>&#34;io&#34;;
    <a id="L81"></a>&#34;os&#34;;
    <a id="L82"></a>&#34;reflect&#34;;
    <a id="L83"></a>&#34;utf8&#34;;
<a id="L84"></a>)

<a id="L86"></a><span class="comment">// State represents the printer state passed to custom formatters.</span>
<a id="L87"></a><span class="comment">// It provides access to the io.Writer interface plus information about</span>
<a id="L88"></a><span class="comment">// the flags and options for the operand&#39;s format specifier.</span>
<a id="L89"></a>type State interface {
    <a id="L90"></a><span class="comment">// Write is the function to call to emit formatted output to be printed.</span>
    <a id="L91"></a>Write(b []byte) (ret int, err os.Error);
    <a id="L92"></a><span class="comment">// Width returns the value of the width option and whether it has been set.</span>
    <a id="L93"></a>Width() (wid int, ok bool);
    <a id="L94"></a><span class="comment">// Precision returns the value of the precision option and whether it has been set.</span>
    <a id="L95"></a>Precision() (prec int, ok bool);

    <a id="L97"></a><span class="comment">// Flag returns whether the flag c, a character, has been set.</span>
    <a id="L98"></a>Flag(int) bool;
<a id="L99"></a>}

<a id="L101"></a><span class="comment">// Formatter is the interface implemented by values with a custom formatter.</span>
<a id="L102"></a><span class="comment">// The implementation of Format may call Sprintf or Fprintf(f) etc.</span>
<a id="L103"></a><span class="comment">// to generate its output.</span>
<a id="L104"></a>type Formatter interface {
    <a id="L105"></a>Format(f State, c int);
<a id="L106"></a>}

<a id="L108"></a><span class="comment">// Stringer is implemented by any value that has a String method(),</span>
<a id="L109"></a><span class="comment">// which defines the ``native&#39;&#39; format for that value.</span>
<a id="L110"></a><span class="comment">// The String method is used to print values passed as an operand</span>
<a id="L111"></a><span class="comment">// to a %s or %v format or to an unformatted printer such as Print.</span>
<a id="L112"></a>type Stringer interface {
    <a id="L113"></a>String() string;
<a id="L114"></a>}

<a id="L116"></a><span class="comment">// GoStringer is implemented by any value that has a GoString() method,</span>
<a id="L117"></a><span class="comment">// which defines the Go syntax for that value.</span>
<a id="L118"></a><span class="comment">// The GoString method is used to print values passed as an operand</span>
<a id="L119"></a><span class="comment">// to a %#v format.</span>
<a id="L120"></a>type GoStringer interface {
    <a id="L121"></a>GoString() string;
<a id="L122"></a>}

<a id="L124"></a>const runeSelf = utf8.RuneSelf
<a id="L125"></a>const allocSize = 32

<a id="L127"></a>type pp struct {
    <a id="L128"></a>n   int;
    <a id="L129"></a>buf []byte;
    <a id="L130"></a>fmt *Fmt;
<a id="L131"></a>}

<a id="L133"></a>func newPrinter() *pp {
    <a id="L134"></a>p := new(pp);
    <a id="L135"></a>p.fmt = New();
    <a id="L136"></a>return p;
<a id="L137"></a>}

<a id="L139"></a>func (p *pp) Width() (wid int, ok bool) { return p.fmt.wid, p.fmt.wid_present }

<a id="L141"></a>func (p *pp) Precision() (prec int, ok bool) { return p.fmt.prec, p.fmt.prec_present }

<a id="L143"></a>func (p *pp) Flag(b int) bool {
    <a id="L144"></a>switch b {
    <a id="L145"></a>case &#39;-&#39;:
        <a id="L146"></a>return p.fmt.minus
    <a id="L147"></a>case &#39;+&#39;:
        <a id="L148"></a>return p.fmt.plus
    <a id="L149"></a>case &#39;#&#39;:
        <a id="L150"></a>return p.fmt.sharp
    <a id="L151"></a>case &#39; &#39;:
        <a id="L152"></a>return p.fmt.space
    <a id="L153"></a>case &#39;0&#39;:
        <a id="L154"></a>return p.fmt.zero
    <a id="L155"></a>}
    <a id="L156"></a>return false;
<a id="L157"></a>}

<a id="L159"></a>func (p *pp) ensure(n int) {
    <a id="L160"></a>if len(p.buf) &lt; n {
        <a id="L161"></a>newn := allocSize + len(p.buf);
        <a id="L162"></a>if newn &lt; n {
            <a id="L163"></a>newn = n + allocSize
        <a id="L164"></a>}
        <a id="L165"></a>b := make([]byte, newn);
        <a id="L166"></a>for i := 0; i &lt; p.n; i++ {
            <a id="L167"></a>b[i] = p.buf[i]
        <a id="L168"></a>}
        <a id="L169"></a>p.buf = b;
    <a id="L170"></a>}
<a id="L171"></a>}

<a id="L173"></a>func (p *pp) addstr(s string) {
    <a id="L174"></a>n := len(s);
    <a id="L175"></a>p.ensure(p.n + n);
    <a id="L176"></a>for i := 0; i &lt; n; i++ {
        <a id="L177"></a>p.buf[p.n] = s[i];
        <a id="L178"></a>p.n++;
    <a id="L179"></a>}
<a id="L180"></a>}

<a id="L182"></a>func (p *pp) addbytes(b []byte, start, end int) {
    <a id="L183"></a>p.ensure(p.n + end - start);
    <a id="L184"></a>for i := start; i &lt; end; i++ {
        <a id="L185"></a>p.buf[p.n] = b[i];
        <a id="L186"></a>p.n++;
    <a id="L187"></a>}
<a id="L188"></a>}

<a id="L190"></a>func (p *pp) add(c int) {
    <a id="L191"></a>p.ensure(p.n + 1);
    <a id="L192"></a>if c &lt; runeSelf {
        <a id="L193"></a>p.buf[p.n] = byte(c);
        <a id="L194"></a>p.n++;
    <a id="L195"></a>} else {
        <a id="L196"></a>p.addstr(string(c))
    <a id="L197"></a>}
<a id="L198"></a>}

<a id="L200"></a><span class="comment">// Implement Write so we can call fprintf on a P, for</span>
<a id="L201"></a><span class="comment">// recursive use in custom verbs.</span>
<a id="L202"></a>func (p *pp) Write(b []byte) (ret int, err os.Error) {
    <a id="L203"></a>p.addbytes(b, 0, len(b));
    <a id="L204"></a>return len(b), nil;
<a id="L205"></a>}

<a id="L207"></a><span class="comment">// These routines end in &#39;f&#39; and take a format string.</span>

<a id="L209"></a><span class="comment">// Fprintf formats according to a format specifier and writes to w.</span>
<a id="L210"></a>func Fprintf(w io.Writer, format string, a ...) (n int, error os.Error) {
    <a id="L211"></a>v := reflect.NewValue(a).(*reflect.StructValue);
    <a id="L212"></a>p := newPrinter();
    <a id="L213"></a>p.doprintf(format, v);
    <a id="L214"></a>n, error = w.Write(p.buf[0:p.n]);
    <a id="L215"></a>return n, error;
<a id="L216"></a>}

<a id="L218"></a><span class="comment">// Printf formats according to a format specifier and writes to standard output.</span>
<a id="L219"></a>func Printf(format string, v ...) (n int, errno os.Error) {
    <a id="L220"></a>n, errno = Fprintf(os.Stdout, format, v);
    <a id="L221"></a>return n, errno;
<a id="L222"></a>}

<a id="L224"></a><span class="comment">// Sprintf formats according to a format specifier and returns the resulting string.</span>
<a id="L225"></a>func Sprintf(format string, a ...) string {
    <a id="L226"></a>v := reflect.NewValue(a).(*reflect.StructValue);
    <a id="L227"></a>p := newPrinter();
    <a id="L228"></a>p.doprintf(format, v);
    <a id="L229"></a>s := string(p.buf)[0:p.n];
    <a id="L230"></a>return s;
<a id="L231"></a>}

<a id="L233"></a><span class="comment">// These routines do not take a format string</span>

<a id="L235"></a><span class="comment">// Fprint formats using the default formats for its operands and writes to w.</span>
<a id="L236"></a><span class="comment">// Spaces are added between operands when neither is a string.</span>
<a id="L237"></a>func Fprint(w io.Writer, a ...) (n int, error os.Error) {
    <a id="L238"></a>v := reflect.NewValue(a).(*reflect.StructValue);
    <a id="L239"></a>p := newPrinter();
    <a id="L240"></a>p.doprint(v, false, false);
    <a id="L241"></a>n, error = w.Write(p.buf[0:p.n]);
    <a id="L242"></a>return n, error;
<a id="L243"></a>}

<a id="L245"></a><span class="comment">// Print formats using the default formats for its operands and writes to standard output.</span>
<a id="L246"></a><span class="comment">// Spaces are added between operands when neither is a string.</span>
<a id="L247"></a>func Print(v ...) (n int, errno os.Error) {
    <a id="L248"></a>n, errno = Fprint(os.Stdout, v);
    <a id="L249"></a>return n, errno;
<a id="L250"></a>}

<a id="L252"></a><span class="comment">// Sprint formats using the default formats for its operands and returns the resulting string.</span>
<a id="L253"></a><span class="comment">// Spaces are added between operands when neither is a string.</span>
<a id="L254"></a>func Sprint(a ...) string {
    <a id="L255"></a>v := reflect.NewValue(a).(*reflect.StructValue);
    <a id="L256"></a>p := newPrinter();
    <a id="L257"></a>p.doprint(v, false, false);
    <a id="L258"></a>s := string(p.buf)[0:p.n];
    <a id="L259"></a>return s;
<a id="L260"></a>}

<a id="L262"></a><span class="comment">// These routines end in &#39;ln&#39;, do not take a format string,</span>
<a id="L263"></a><span class="comment">// always add spaces between operands, and add a newline</span>
<a id="L264"></a><span class="comment">// after the last operand.</span>

<a id="L266"></a><span class="comment">// Fprintln formats using the default formats for its operands and writes to w.</span>
<a id="L267"></a><span class="comment">// Spaces are always added between operands and a newline is appended.</span>
<a id="L268"></a>func Fprintln(w io.Writer, a ...) (n int, error os.Error) {
    <a id="L269"></a>v := reflect.NewValue(a).(*reflect.StructValue);
    <a id="L270"></a>p := newPrinter();
    <a id="L271"></a>p.doprint(v, true, true);
    <a id="L272"></a>n, error = w.Write(p.buf[0:p.n]);
    <a id="L273"></a>return n, error;
<a id="L274"></a>}

<a id="L276"></a><span class="comment">// Println formats using the default formats for its operands and writes to standard output.</span>
<a id="L277"></a><span class="comment">// Spaces are always added between operands and a newline is appended.</span>
<a id="L278"></a>func Println(v ...) (n int, errno os.Error) {
    <a id="L279"></a>n, errno = Fprintln(os.Stdout, v);
    <a id="L280"></a>return n, errno;
<a id="L281"></a>}

<a id="L283"></a><span class="comment">// Sprintln formats using the default formats for its operands and returns the resulting string.</span>
<a id="L284"></a><span class="comment">// Spaces are always added between operands and a newline is appended.</span>
<a id="L285"></a>func Sprintln(a ...) string {
    <a id="L286"></a>v := reflect.NewValue(a).(*reflect.StructValue);
    <a id="L287"></a>p := newPrinter();
    <a id="L288"></a>p.doprint(v, true, true);
    <a id="L289"></a>s := string(p.buf)[0:p.n];
    <a id="L290"></a>return s;
<a id="L291"></a>}


<a id="L294"></a><span class="comment">// Get the i&#39;th arg of the struct value.</span>
<a id="L295"></a><span class="comment">// If the arg itself is an interface, return a value for</span>
<a id="L296"></a><span class="comment">// the thing inside the interface, not the interface itself.</span>
<a id="L297"></a>func getField(v *reflect.StructValue, i int) reflect.Value {
    <a id="L298"></a>val := v.Field(i);
    <a id="L299"></a>if i, ok := val.(*reflect.InterfaceValue); ok {
        <a id="L300"></a>if inter := i.Interface(); inter != nil {
            <a id="L301"></a>return reflect.NewValue(inter)
        <a id="L302"></a>}
    <a id="L303"></a>}
    <a id="L304"></a>return val;
<a id="L305"></a>}

<a id="L307"></a><span class="comment">// Getters for the fields of the argument structure.</span>

<a id="L309"></a>func getBool(v reflect.Value) (val bool, ok bool) {
    <a id="L310"></a>if b, ok := v.(*reflect.BoolValue); ok {
        <a id="L311"></a>return b.Get(), true
    <a id="L312"></a>}
    <a id="L313"></a>return;
<a id="L314"></a>}

<a id="L316"></a>func getInt(v reflect.Value) (val int64, signed, ok bool) {
    <a id="L317"></a>switch v := v.(type) {
    <a id="L318"></a>case *reflect.IntValue:
        <a id="L319"></a>return int64(v.Get()), true, true
    <a id="L320"></a>case *reflect.Int8Value:
        <a id="L321"></a>return int64(v.Get()), true, true
    <a id="L322"></a>case *reflect.Int16Value:
        <a id="L323"></a>return int64(v.Get()), true, true
    <a id="L324"></a>case *reflect.Int32Value:
        <a id="L325"></a>return int64(v.Get()), true, true
    <a id="L326"></a>case *reflect.Int64Value:
        <a id="L327"></a>return int64(v.Get()), true, true
    <a id="L328"></a>case *reflect.UintValue:
        <a id="L329"></a>return int64(v.Get()), false, true
    <a id="L330"></a>case *reflect.Uint8Value:
        <a id="L331"></a>return int64(v.Get()), false, true
    <a id="L332"></a>case *reflect.Uint16Value:
        <a id="L333"></a>return int64(v.Get()), false, true
    <a id="L334"></a>case *reflect.Uint32Value:
        <a id="L335"></a>return int64(v.Get()), false, true
    <a id="L336"></a>case *reflect.Uint64Value:
        <a id="L337"></a>return int64(v.Get()), false, true
    <a id="L338"></a>case *reflect.UintptrValue:
        <a id="L339"></a>return int64(v.Get()), false, true
    <a id="L340"></a>}
    <a id="L341"></a>return;
<a id="L342"></a>}

<a id="L344"></a>func getString(v reflect.Value) (val string, ok bool) {
    <a id="L345"></a>if v, ok := v.(*reflect.StringValue); ok {
        <a id="L346"></a>return v.Get(), true
    <a id="L347"></a>}
    <a id="L348"></a>if bytes, ok := v.Interface().([]byte); ok {
        <a id="L349"></a>return string(bytes), true
    <a id="L350"></a>}
    <a id="L351"></a>return;
<a id="L352"></a>}

<a id="L354"></a>func getFloat32(v reflect.Value) (val float32, ok bool) {
    <a id="L355"></a>switch v := v.(type) {
    <a id="L356"></a>case *reflect.Float32Value:
        <a id="L357"></a>return float32(v.Get()), true
    <a id="L358"></a>case *reflect.FloatValue:
        <a id="L359"></a>if v.Type().Size()*8 == 32 {
            <a id="L360"></a>return float32(v.Get()), true
        <a id="L361"></a>}
    <a id="L362"></a>}
    <a id="L363"></a>return;
<a id="L364"></a>}

<a id="L366"></a>func getFloat64(v reflect.Value) (val float64, ok bool) {
    <a id="L367"></a>switch v := v.(type) {
    <a id="L368"></a>case *reflect.FloatValue:
        <a id="L369"></a>if v.Type().Size()*8 == 64 {
            <a id="L370"></a>return float64(v.Get()), true
        <a id="L371"></a>}
    <a id="L372"></a>case *reflect.Float64Value:
        <a id="L373"></a>return float64(v.Get()), true
    <a id="L374"></a>}
    <a id="L375"></a>return;
<a id="L376"></a>}

<a id="L378"></a>func getPtr(v reflect.Value) (val uintptr, ok bool) {
    <a id="L379"></a>switch v := v.(type) {
    <a id="L380"></a>case *reflect.PtrValue:
        <a id="L381"></a>return uintptr(v.Get()), true
    <a id="L382"></a>}
    <a id="L383"></a>return;
<a id="L384"></a>}

<a id="L386"></a><span class="comment">// Convert ASCII to integer.  n is 0 (and got is false) if no number present.</span>

<a id="L388"></a>func parsenum(s string, start, end int) (n int, got bool, newi int) {
    <a id="L389"></a>if start &gt;= end {
        <a id="L390"></a>return 0, false, end
    <a id="L391"></a>}
    <a id="L392"></a>isnum := false;
    <a id="L393"></a>num := 0;
    <a id="L394"></a>for &#39;0&#39; &lt;= s[start] &amp;&amp; s[start] &lt;= &#39;9&#39; {
        <a id="L395"></a>num = num*10 + int(s[start]-&#39;0&#39;);
        <a id="L396"></a>start++;
        <a id="L397"></a>isnum = true;
    <a id="L398"></a>}
    <a id="L399"></a>return num, isnum, start;
<a id="L400"></a>}

<a id="L402"></a>type uintptrGetter interface {
    <a id="L403"></a>Get() uintptr;
<a id="L404"></a>}

<a id="L406"></a>func (p *pp) printField(field reflect.Value, plus, sharp bool, depth int) (was_string bool) {
    <a id="L407"></a>inter := field.Interface();
    <a id="L408"></a>if inter != nil {
        <a id="L409"></a>switch {
        <a id="L410"></a>default:
            <a id="L411"></a>if stringer, ok := inter.(Stringer); ok {
                <a id="L412"></a>p.addstr(stringer.String());
                <a id="L413"></a>return false; <span class="comment">// this value is not a string</span>
            <a id="L414"></a>}
        <a id="L415"></a>case sharp:
            <a id="L416"></a>if stringer, ok := inter.(GoStringer); ok {
                <a id="L417"></a>p.addstr(stringer.GoString());
                <a id="L418"></a>return false; <span class="comment">// this value is not a string</span>
            <a id="L419"></a>}
        <a id="L420"></a>}
    <a id="L421"></a>}
    <a id="L422"></a>s := &#34;&#34;;
<a id="L423"></a>BigSwitch:
    <a id="L424"></a>switch f := field.(type) {
    <a id="L425"></a>case *reflect.BoolValue:
        <a id="L426"></a>s = p.fmt.Fmt_boolean(f.Get()).Str()
    <a id="L427"></a>case *reflect.Float32Value:
        <a id="L428"></a>s = p.fmt.Fmt_g32(f.Get()).Str()
    <a id="L429"></a>case *reflect.Float64Value:
        <a id="L430"></a>s = p.fmt.Fmt_g64(f.Get()).Str()
    <a id="L431"></a>case *reflect.FloatValue:
        <a id="L432"></a>if field.Type().Size()*8 == 32 {
            <a id="L433"></a>s = p.fmt.Fmt_g32(float32(f.Get())).Str()
        <a id="L434"></a>} else {
            <a id="L435"></a>s = p.fmt.Fmt_g64(float64(f.Get())).Str()
        <a id="L436"></a>}
    <a id="L437"></a>case *reflect.StringValue:
        <a id="L438"></a>if sharp {
            <a id="L439"></a>s = p.fmt.Fmt_q(f.Get()).Str()
        <a id="L440"></a>} else {
            <a id="L441"></a>s = p.fmt.Fmt_s(f.Get()).Str();
            <a id="L442"></a>was_string = true;
        <a id="L443"></a>}
    <a id="L444"></a>case *reflect.MapValue:
        <a id="L445"></a>if sharp {
            <a id="L446"></a>p.addstr(field.Type().String());
            <a id="L447"></a>p.addstr(&#34;{&#34;);
        <a id="L448"></a>} else {
            <a id="L449"></a>p.addstr(&#34;map[&#34;)
        <a id="L450"></a>}
        <a id="L451"></a>keys := f.Keys();
        <a id="L452"></a>for i, key := range keys {
            <a id="L453"></a>if i &gt; 0 {
                <a id="L454"></a>if sharp {
                    <a id="L455"></a>p.addstr(&#34;, &#34;)
                <a id="L456"></a>} else {
                    <a id="L457"></a>p.addstr(&#34; &#34;)
                <a id="L458"></a>}
            <a id="L459"></a>}
            <a id="L460"></a>p.printField(key, plus, sharp, depth+1);
            <a id="L461"></a>p.addstr(&#34;:&#34;);
            <a id="L462"></a>p.printField(f.Elem(key), plus, sharp, depth+1);
        <a id="L463"></a>}
        <a id="L464"></a>if sharp {
            <a id="L465"></a>p.addstr(&#34;}&#34;)
        <a id="L466"></a>} else {
            <a id="L467"></a>p.addstr(&#34;]&#34;)
        <a id="L468"></a>}
    <a id="L469"></a>case *reflect.StructValue:
        <a id="L470"></a>if sharp {
            <a id="L471"></a>p.addstr(field.Type().String())
        <a id="L472"></a>}
        <a id="L473"></a>p.add(&#39;{&#39;);
        <a id="L474"></a>v := f;
        <a id="L475"></a>t := v.Type().(*reflect.StructType);
        <a id="L476"></a>p.fmt.clearflags(); <span class="comment">// clear flags for p.printField</span>
        <a id="L477"></a>for i := 0; i &lt; v.NumField(); i++ {
            <a id="L478"></a>if i &gt; 0 {
                <a id="L479"></a>if sharp {
                    <a id="L480"></a>p.addstr(&#34;, &#34;)
                <a id="L481"></a>} else {
                    <a id="L482"></a>p.addstr(&#34; &#34;)
                <a id="L483"></a>}
            <a id="L484"></a>}
            <a id="L485"></a>if plus || sharp {
                <a id="L486"></a>if f := t.Field(i); f.Name != &#34;&#34; {
                    <a id="L487"></a>p.addstr(f.Name);
                    <a id="L488"></a>p.add(&#39;:&#39;);
                <a id="L489"></a>}
            <a id="L490"></a>}
            <a id="L491"></a>p.printField(getField(v, i), plus, sharp, depth+1);
        <a id="L492"></a>}
        <a id="L493"></a>p.addstr(&#34;}&#34;);
    <a id="L494"></a>case *reflect.InterfaceValue:
        <a id="L495"></a>value := f.Elem();
        <a id="L496"></a>if value == nil {
            <a id="L497"></a>if sharp {
                <a id="L498"></a>p.addstr(field.Type().String());
                <a id="L499"></a>p.addstr(&#34;(nil)&#34;);
            <a id="L500"></a>} else {
                <a id="L501"></a>s = &#34;&lt;nil&gt;&#34;
            <a id="L502"></a>}
        <a id="L503"></a>} else {
            <a id="L504"></a>return p.printField(value, plus, sharp, depth+1)
        <a id="L505"></a>}
    <a id="L506"></a>case reflect.ArrayOrSliceValue:
        <a id="L507"></a>if sharp {
            <a id="L508"></a>p.addstr(field.Type().String());
            <a id="L509"></a>p.addstr(&#34;{&#34;);
        <a id="L510"></a>} else {
            <a id="L511"></a>p.addstr(&#34;[&#34;)
        <a id="L512"></a>}
        <a id="L513"></a>for i := 0; i &lt; f.Len(); i++ {
            <a id="L514"></a>if i &gt; 0 {
                <a id="L515"></a>if sharp {
                    <a id="L516"></a>p.addstr(&#34;, &#34;)
                <a id="L517"></a>} else {
                    <a id="L518"></a>p.addstr(&#34; &#34;)
                <a id="L519"></a>}
            <a id="L520"></a>}
            <a id="L521"></a>p.printField(f.Elem(i), plus, sharp, depth+1);
        <a id="L522"></a>}
        <a id="L523"></a>if sharp {
            <a id="L524"></a>p.addstr(&#34;}&#34;)
        <a id="L525"></a>} else {
            <a id="L526"></a>p.addstr(&#34;]&#34;)
        <a id="L527"></a>}
    <a id="L528"></a>case *reflect.PtrValue:
        <a id="L529"></a>v := f.Get();
        <a id="L530"></a><span class="comment">// pointer to array or slice or struct?  ok at top level</span>
        <a id="L531"></a><span class="comment">// but not embedded (avoid loops)</span>
        <a id="L532"></a>if v != 0 &amp;&amp; depth == 0 {
            <a id="L533"></a>switch a := f.Elem().(type) {
            <a id="L534"></a>case reflect.ArrayOrSliceValue:
                <a id="L535"></a>p.addstr(&#34;&amp;&#34;);
                <a id="L536"></a>p.printField(a, plus, sharp, depth+1);
                <a id="L537"></a>break BigSwitch;
            <a id="L538"></a>case *reflect.StructValue:
                <a id="L539"></a>p.addstr(&#34;&amp;&#34;);
                <a id="L540"></a>p.printField(a, plus, sharp, depth+1);
                <a id="L541"></a>break BigSwitch;
            <a id="L542"></a>}
        <a id="L543"></a>}
        <a id="L544"></a>if sharp {
            <a id="L545"></a>p.addstr(&#34;(&#34;);
            <a id="L546"></a>p.addstr(field.Type().String());
            <a id="L547"></a>p.addstr(&#34;)(&#34;);
            <a id="L548"></a>if v == 0 {
                <a id="L549"></a>p.addstr(&#34;nil&#34;)
            <a id="L550"></a>} else {
                <a id="L551"></a>p.fmt.sharp = true;
                <a id="L552"></a>p.addstr(p.fmt.Fmt_ux64(uint64(v)).Str());
            <a id="L553"></a>}
            <a id="L554"></a>p.addstr(&#34;)&#34;);
            <a id="L555"></a>break;
        <a id="L556"></a>}
        <a id="L557"></a>if v == 0 {
            <a id="L558"></a>s = &#34;&lt;nil&gt;&#34;;
            <a id="L559"></a>break;
        <a id="L560"></a>}
        <a id="L561"></a>p.fmt.sharp = true; <span class="comment">// turn 0x on</span>
        <a id="L562"></a>s = p.fmt.Fmt_ux64(uint64(v)).Str();
    <a id="L563"></a>case uintptrGetter:
        <a id="L564"></a>v := f.Get();
        <a id="L565"></a>if sharp {
            <a id="L566"></a>p.addstr(&#34;(&#34;);
            <a id="L567"></a>p.addstr(field.Type().String());
            <a id="L568"></a>p.addstr(&#34;)(&#34;);
            <a id="L569"></a>if v == 0 {
                <a id="L570"></a>p.addstr(&#34;nil&#34;)
            <a id="L571"></a>} else {
                <a id="L572"></a>p.fmt.sharp = true;
                <a id="L573"></a>p.addstr(p.fmt.Fmt_ux64(uint64(v)).Str());
            <a id="L574"></a>}
            <a id="L575"></a>p.addstr(&#34;)&#34;);
        <a id="L576"></a>} else {
            <a id="L577"></a>p.fmt.sharp = true; <span class="comment">// turn 0x on</span>
            <a id="L578"></a>p.addstr(p.fmt.Fmt_ux64(uint64(f.Get())).Str());
        <a id="L579"></a>}
    <a id="L580"></a>default:
        <a id="L581"></a>v, signed, ok := getInt(field);
        <a id="L582"></a>if ok {
            <a id="L583"></a>if signed {
                <a id="L584"></a>s = p.fmt.Fmt_d64(v).Str()
            <a id="L585"></a>} else {
                <a id="L586"></a>if sharp {
                    <a id="L587"></a>p.fmt.sharp = true; <span class="comment">// turn on 0x</span>
                    <a id="L588"></a>s = p.fmt.Fmt_ux64(uint64(v)).Str();
                <a id="L589"></a>} else {
                    <a id="L590"></a>s = p.fmt.Fmt_ud64(uint64(v)).Str()
                <a id="L591"></a>}
            <a id="L592"></a>}
            <a id="L593"></a>break;
        <a id="L594"></a>}
        <a id="L595"></a>s = &#34;?&#34; + field.Type().String() + &#34;?&#34;;
    <a id="L596"></a>}
    <a id="L597"></a>p.addstr(s);
    <a id="L598"></a>return was_string;
<a id="L599"></a>}

<a id="L601"></a>func (p *pp) doprintf(format string, v *reflect.StructValue) {
    <a id="L602"></a>p.ensure(len(format)); <span class="comment">// a good starting size</span>
    <a id="L603"></a>end := len(format) - 1;
    <a id="L604"></a>fieldnum := 0; <span class="comment">// we process one field per non-trivial format</span>
    <a id="L605"></a>for i := 0; i &lt;= end; {
        <a id="L606"></a>c, w := utf8.DecodeRuneInString(format[i:len(format)]);
        <a id="L607"></a>if c != &#39;%&#39; || i == end {
            <a id="L608"></a>p.add(c);
            <a id="L609"></a>i += w;
            <a id="L610"></a>continue;
        <a id="L611"></a>}
        <a id="L612"></a>i++;
        <a id="L613"></a><span class="comment">// flags and widths</span>
        <a id="L614"></a>p.fmt.clearflags();
    <a id="L615"></a>F:  for ; i &lt; end; i++ {
            <a id="L616"></a>switch format[i] {
            <a id="L617"></a>case &#39;#&#39;:
                <a id="L618"></a>p.fmt.sharp = true
            <a id="L619"></a>case &#39;0&#39;:
                <a id="L620"></a>p.fmt.zero = true
            <a id="L621"></a>case &#39;+&#39;:
                <a id="L622"></a>p.fmt.plus = true
            <a id="L623"></a>case &#39;-&#39;:
                <a id="L624"></a>p.fmt.minus = true
            <a id="L625"></a>case &#39; &#39;:
                <a id="L626"></a>p.fmt.space = true
            <a id="L627"></a>default:
                <a id="L628"></a>break F
            <a id="L629"></a>}
        <a id="L630"></a>}
        <a id="L631"></a><span class="comment">// do we have 20 (width)?</span>
        <a id="L632"></a>p.fmt.wid, p.fmt.wid_present, i = parsenum(format, i, end);
        <a id="L633"></a><span class="comment">// do we have .20 (precision)?</span>
        <a id="L634"></a>if i &lt; end &amp;&amp; format[i] == &#39;.&#39; {
            <a id="L635"></a>p.fmt.prec, p.fmt.prec_present, i = parsenum(format, i+1, end)
        <a id="L636"></a>}
        <a id="L637"></a>c, w = utf8.DecodeRuneInString(format[i:len(format)]);
        <a id="L638"></a>i += w;
        <a id="L639"></a><span class="comment">// percent is special - absorbs no operand</span>
        <a id="L640"></a>if c == &#39;%&#39; {
            <a id="L641"></a>p.add(&#39;%&#39;); <span class="comment">// TODO: should we bother with width &amp; prec?</span>
            <a id="L642"></a>continue;
        <a id="L643"></a>}
        <a id="L644"></a>if fieldnum &gt;= v.NumField() { <span class="comment">// out of operands</span>
            <a id="L645"></a>p.add(&#39;%&#39;);
            <a id="L646"></a>p.add(c);
            <a id="L647"></a>p.addstr(&#34;(missing)&#34;);
            <a id="L648"></a>continue;
        <a id="L649"></a>}
        <a id="L650"></a>field := getField(v, fieldnum);
        <a id="L651"></a>fieldnum++;

        <a id="L653"></a><span class="comment">// Try formatter except for %T,</span>
        <a id="L654"></a><span class="comment">// which is special and handled internally.</span>
        <a id="L655"></a>inter := field.Interface();
        <a id="L656"></a>if inter != nil &amp;&amp; c != &#39;T&#39; {
            <a id="L657"></a>if formatter, ok := inter.(Formatter); ok {
                <a id="L658"></a>formatter.Format(p, c);
                <a id="L659"></a>continue;
            <a id="L660"></a>}
        <a id="L661"></a>}

        <a id="L663"></a>s := &#34;&#34;;
        <a id="L664"></a>switch c {
        <a id="L665"></a><span class="comment">// bool</span>
        <a id="L666"></a>case &#39;t&#39;:
            <a id="L667"></a>if v, ok := getBool(field); ok {
                <a id="L668"></a>if v {
                    <a id="L669"></a>s = &#34;true&#34;
                <a id="L670"></a>} else {
                    <a id="L671"></a>s = &#34;false&#34;
                <a id="L672"></a>}
            <a id="L673"></a>} else {
                <a id="L674"></a>goto badtype
            <a id="L675"></a>}

        <a id="L677"></a><span class="comment">// int</span>
        <a id="L678"></a>case &#39;b&#39;:
            <a id="L679"></a>if v, _, ok := getInt(field); ok {
                <a id="L680"></a>s = p.fmt.Fmt_b64(uint64(v)).Str() <span class="comment">// always unsigned</span>
            <a id="L681"></a>} else if v, ok := getFloat32(field); ok {
                <a id="L682"></a>s = p.fmt.Fmt_fb32(v).Str()
            <a id="L683"></a>} else if v, ok := getFloat64(field); ok {
                <a id="L684"></a>s = p.fmt.Fmt_fb64(v).Str()
            <a id="L685"></a>} else {
                <a id="L686"></a>goto badtype
            <a id="L687"></a>}
        <a id="L688"></a>case &#39;c&#39;:
            <a id="L689"></a>if v, _, ok := getInt(field); ok {
                <a id="L690"></a>s = p.fmt.Fmt_c(int(v)).Str()
            <a id="L691"></a>} else {
                <a id="L692"></a>goto badtype
            <a id="L693"></a>}
        <a id="L694"></a>case &#39;d&#39;:
            <a id="L695"></a>if v, signed, ok := getInt(field); ok {
                <a id="L696"></a>if signed {
                    <a id="L697"></a>s = p.fmt.Fmt_d64(v).Str()
                <a id="L698"></a>} else {
                    <a id="L699"></a>s = p.fmt.Fmt_ud64(uint64(v)).Str()
                <a id="L700"></a>}
            <a id="L701"></a>} else {
                <a id="L702"></a>goto badtype
            <a id="L703"></a>}
        <a id="L704"></a>case &#39;o&#39;:
            <a id="L705"></a>if v, signed, ok := getInt(field); ok {
                <a id="L706"></a>if signed {
                    <a id="L707"></a>s = p.fmt.Fmt_o64(v).Str()
                <a id="L708"></a>} else {
                    <a id="L709"></a>s = p.fmt.Fmt_uo64(uint64(v)).Str()
                <a id="L710"></a>}
            <a id="L711"></a>} else {
                <a id="L712"></a>goto badtype
            <a id="L713"></a>}
        <a id="L714"></a>case &#39;x&#39;:
            <a id="L715"></a>if v, signed, ok := getInt(field); ok {
                <a id="L716"></a>if signed {
                    <a id="L717"></a>s = p.fmt.Fmt_x64(v).Str()
                <a id="L718"></a>} else {
                    <a id="L719"></a>s = p.fmt.Fmt_ux64(uint64(v)).Str()
                <a id="L720"></a>}
            <a id="L721"></a>} else if v, ok := getString(field); ok {
                <a id="L722"></a>s = p.fmt.Fmt_sx(v).Str()
            <a id="L723"></a>} else {
                <a id="L724"></a>goto badtype
            <a id="L725"></a>}
        <a id="L726"></a>case &#39;X&#39;:
            <a id="L727"></a>if v, signed, ok := getInt(field); ok {
                <a id="L728"></a>if signed {
                    <a id="L729"></a>s = p.fmt.Fmt_X64(v).Str()
                <a id="L730"></a>} else {
                    <a id="L731"></a>s = p.fmt.Fmt_uX64(uint64(v)).Str()
                <a id="L732"></a>}
            <a id="L733"></a>} else if v, ok := getString(field); ok {
                <a id="L734"></a>s = p.fmt.Fmt_sX(v).Str()
            <a id="L735"></a>} else {
                <a id="L736"></a>goto badtype
            <a id="L737"></a>}

        <a id="L739"></a><span class="comment">// float</span>
        <a id="L740"></a>case &#39;e&#39;:
            <a id="L741"></a>if v, ok := getFloat32(field); ok {
                <a id="L742"></a>s = p.fmt.Fmt_e32(v).Str()
            <a id="L743"></a>} else if v, ok := getFloat64(field); ok {
                <a id="L744"></a>s = p.fmt.Fmt_e64(v).Str()
            <a id="L745"></a>} else {
                <a id="L746"></a>goto badtype
            <a id="L747"></a>}
        <a id="L748"></a>case &#39;E&#39;:
            <a id="L749"></a>if v, ok := getFloat32(field); ok {
                <a id="L750"></a>s = p.fmt.Fmt_E32(v).Str()
            <a id="L751"></a>} else if v, ok := getFloat64(field); ok {
                <a id="L752"></a>s = p.fmt.Fmt_E64(v).Str()
            <a id="L753"></a>} else {
                <a id="L754"></a>goto badtype
            <a id="L755"></a>}
        <a id="L756"></a>case &#39;f&#39;:
            <a id="L757"></a>if v, ok := getFloat32(field); ok {
                <a id="L758"></a>s = p.fmt.Fmt_f32(v).Str()
            <a id="L759"></a>} else if v, ok := getFloat64(field); ok {
                <a id="L760"></a>s = p.fmt.Fmt_f64(v).Str()
            <a id="L761"></a>} else {
                <a id="L762"></a>goto badtype
            <a id="L763"></a>}
        <a id="L764"></a>case &#39;g&#39;:
            <a id="L765"></a>if v, ok := getFloat32(field); ok {
                <a id="L766"></a>s = p.fmt.Fmt_g32(v).Str()
            <a id="L767"></a>} else if v, ok := getFloat64(field); ok {
                <a id="L768"></a>s = p.fmt.Fmt_g64(v).Str()
            <a id="L769"></a>} else {
                <a id="L770"></a>goto badtype
            <a id="L771"></a>}
        <a id="L772"></a>case &#39;G&#39;:
            <a id="L773"></a>if v, ok := getFloat32(field); ok {
                <a id="L774"></a>s = p.fmt.Fmt_G32(v).Str()
            <a id="L775"></a>} else if v, ok := getFloat64(field); ok {
                <a id="L776"></a>s = p.fmt.Fmt_G64(v).Str()
            <a id="L777"></a>} else {
                <a id="L778"></a>goto badtype
            <a id="L779"></a>}

        <a id="L781"></a><span class="comment">// string</span>
        <a id="L782"></a>case &#39;s&#39;:
            <a id="L783"></a>if inter != nil {
                <a id="L784"></a><span class="comment">// if object implements String, use the result.</span>
                <a id="L785"></a>if stringer, ok := inter.(Stringer); ok {
                    <a id="L786"></a>s = p.fmt.Fmt_s(stringer.String()).Str();
                    <a id="L787"></a>break;
                <a id="L788"></a>}
            <a id="L789"></a>}
            <a id="L790"></a>if v, ok := getString(field); ok {
                <a id="L791"></a>s = p.fmt.Fmt_s(v).Str()
            <a id="L792"></a>} else {
                <a id="L793"></a>goto badtype
            <a id="L794"></a>}
        <a id="L795"></a>case &#39;q&#39;:
            <a id="L796"></a>if v, ok := getString(field); ok {
                <a id="L797"></a>s = p.fmt.Fmt_q(v).Str()
            <a id="L798"></a>} else {
                <a id="L799"></a>goto badtype
            <a id="L800"></a>}

        <a id="L802"></a><span class="comment">// pointer</span>
        <a id="L803"></a>case &#39;p&#39;:
            <a id="L804"></a>if v, ok := getPtr(field); ok {
                <a id="L805"></a>if v == 0 {
                    <a id="L806"></a>s = &#34;&lt;nil&gt;&#34;
                <a id="L807"></a>} else {
                    <a id="L808"></a>s = &#34;0x&#34; + p.fmt.Fmt_uX64(uint64(v)).Str()
                <a id="L809"></a>}
            <a id="L810"></a>} else {
                <a id="L811"></a>goto badtype
            <a id="L812"></a>}

        <a id="L814"></a><span class="comment">// arbitrary value; do your best</span>
        <a id="L815"></a>case &#39;v&#39;:
            <a id="L816"></a>plus, sharp := p.fmt.plus, p.fmt.sharp;
            <a id="L817"></a>p.fmt.plus = false;
            <a id="L818"></a>p.fmt.sharp = false;
            <a id="L819"></a>p.printField(field, plus, sharp, 0);

        <a id="L821"></a><span class="comment">// the value&#39;s type</span>
        <a id="L822"></a>case &#39;T&#39;:
            <a id="L823"></a>s = field.Type().String()

        <a id="L825"></a>default:
        <a id="L826"></a>badtype:
            <a id="L827"></a>s = &#34;%&#34; + string(c) + &#34;(&#34; + field.Type().String() + &#34;=&#34;;
            <a id="L828"></a>p.addstr(s);
            <a id="L829"></a>p.printField(field, false, false, 0);
            <a id="L830"></a>s = &#34;)&#34;;
        <a id="L831"></a>}
        <a id="L832"></a>p.addstr(s);
    <a id="L833"></a>}
    <a id="L834"></a>if fieldnum &lt; v.NumField() {
        <a id="L835"></a>p.addstr(&#34;?(extra &#34;);
        <a id="L836"></a>for ; fieldnum &lt; v.NumField(); fieldnum++ {
            <a id="L837"></a>field := getField(v, fieldnum);
            <a id="L838"></a>p.addstr(field.Type().String());
            <a id="L839"></a>p.addstr(&#34;=&#34;);
            <a id="L840"></a>p.printField(field, false, false, 0);
            <a id="L841"></a>if fieldnum+1 &lt; v.NumField() {
                <a id="L842"></a>p.addstr(&#34;, &#34;)
            <a id="L843"></a>}
        <a id="L844"></a>}
        <a id="L845"></a>p.addstr(&#34;)&#34;);
    <a id="L846"></a>}
<a id="L847"></a>}

<a id="L849"></a>func (p *pp) doprint(v *reflect.StructValue, addspace, addnewline bool) {
    <a id="L850"></a>prev_string := false;
    <a id="L851"></a>for fieldnum := 0; fieldnum &lt; v.NumField(); fieldnum++ {
        <a id="L852"></a><span class="comment">// always add spaces if we&#39;re doing println</span>
        <a id="L853"></a>field := getField(v, fieldnum);
        <a id="L854"></a>if fieldnum &gt; 0 {
            <a id="L855"></a>_, is_string := field.(*reflect.StringValue);
            <a id="L856"></a>if addspace || !is_string &amp;&amp; !prev_string {
                <a id="L857"></a>p.add(&#39; &#39;)
            <a id="L858"></a>}
        <a id="L859"></a>}
        <a id="L860"></a>prev_string = p.printField(field, false, false, 0);
    <a id="L861"></a>}
    <a id="L862"></a>if addnewline {
        <a id="L863"></a>p.add(&#39;\n&#39;)
    <a id="L864"></a>}
<a id="L865"></a>}
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
