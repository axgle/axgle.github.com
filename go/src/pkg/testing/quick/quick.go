<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/testing/quick/quick.go</title>

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
	<li>Thu Nov 12 16:00:00 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/testing/quick/quick.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package implements utility functions to help with black box testing.</span>
<a id="L6"></a>package quick

<a id="L8"></a>import (
    <a id="L9"></a>&#34;flag&#34;;
    <a id="L10"></a>&#34;fmt&#34;;
    <a id="L11"></a>&#34;math&#34;;
    <a id="L12"></a>&#34;os&#34;;
    <a id="L13"></a>&#34;rand&#34;;
    <a id="L14"></a>&#34;reflect&#34;;
    <a id="L15"></a>&#34;strings&#34;;
<a id="L16"></a>)

<a id="L18"></a>var defaultMaxCount *int = flag.Int(&#34;quickchecks&#34;, 100, &#34;The default number of iterations for each check&#34;)

<a id="L20"></a><span class="comment">// A Generator can generate random values of its own type.</span>
<a id="L21"></a>type Generator interface {
    <a id="L22"></a><span class="comment">// Generate returns a random instance of the type on which it is a</span>
    <a id="L23"></a><span class="comment">// method using the size as a size hint.</span>
    <a id="L24"></a>Generate(rand *rand.Rand, size int) reflect.Value;
<a id="L25"></a>}

<a id="L27"></a><span class="comment">// randFloat32 generates a random float taking the full range of a float32.</span>
<a id="L28"></a>func randFloat32(rand *rand.Rand) float32 {
    <a id="L29"></a>f := rand.Float64() * math.MaxFloat32;
    <a id="L30"></a>if rand.Int()&amp;1 == 1 {
        <a id="L31"></a>f = -f
    <a id="L32"></a>}
    <a id="L33"></a>return float32(f);
<a id="L34"></a>}

<a id="L36"></a><span class="comment">// randFloat64 generates a random float taking the full range of a float64.</span>
<a id="L37"></a>func randFloat64(rand *rand.Rand) float64 {
    <a id="L38"></a>f := rand.Float64();
    <a id="L39"></a>if rand.Int()&amp;1 == 1 {
        <a id="L40"></a>f = -f
    <a id="L41"></a>}
    <a id="L42"></a>return f;
<a id="L43"></a>}

<a id="L45"></a><span class="comment">// randInt64 returns a random integer taking half the range of an int64.</span>
<a id="L46"></a>func randInt64(rand *rand.Rand) int64 { return rand.Int63() - 1&lt;&lt;62 }

<a id="L48"></a><span class="comment">// complexSize is the maximum length of arbitrary values that contain other</span>
<a id="L49"></a><span class="comment">// values.</span>
<a id="L50"></a>const complexSize = 50

<a id="L52"></a><span class="comment">// Value returns an arbitrary value of the given type.</span>
<a id="L53"></a><span class="comment">// If the type implements the Generator interface, that will be used.</span>
<a id="L54"></a><span class="comment">// Note: in order to create arbitrary values for structs, all the members must be public.</span>
<a id="L55"></a>func Value(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool) {
    <a id="L56"></a>if m, ok := reflect.MakeZero(t).Interface().(Generator); ok {
        <a id="L57"></a>return m.Generate(rand, complexSize), true
    <a id="L58"></a>}

    <a id="L60"></a>switch concrete := t.(type) {
    <a id="L61"></a>case *reflect.BoolType:
        <a id="L62"></a>return reflect.NewValue(rand.Int()&amp;1 == 0), true
    <a id="L63"></a>case *reflect.Float32Type:
        <a id="L64"></a>return reflect.NewValue(randFloat32(rand)), true
    <a id="L65"></a>case *reflect.Float64Type:
        <a id="L66"></a>return reflect.NewValue(randFloat64(rand)), true
    <a id="L67"></a>case *reflect.FloatType:
        <a id="L68"></a>if t.Size() == 4 {
            <a id="L69"></a>return reflect.NewValue(float(randFloat32(rand))), true
        <a id="L70"></a>} else {
            <a id="L71"></a>return reflect.NewValue(float(randFloat64(rand))), true
        <a id="L72"></a>}
    <a id="L73"></a>case *reflect.Int16Type:
        <a id="L74"></a>return reflect.NewValue(int16(randInt64(rand))), true
    <a id="L75"></a>case *reflect.Int32Type:
        <a id="L76"></a>return reflect.NewValue(int32(randInt64(rand))), true
    <a id="L77"></a>case *reflect.Int64Type:
        <a id="L78"></a>return reflect.NewValue(randInt64(rand)), true
    <a id="L79"></a>case *reflect.Int8Type:
        <a id="L80"></a>return reflect.NewValue(int8(randInt64(rand))), true
    <a id="L81"></a>case *reflect.IntType:
        <a id="L82"></a>return reflect.NewValue(int(randInt64(rand))), true
    <a id="L83"></a>case *reflect.MapType:
        <a id="L84"></a>numElems := rand.Intn(complexSize);
        <a id="L85"></a>m := reflect.MakeMap(concrete);
        <a id="L86"></a>for i := 0; i &lt; numElems; i++ {
            <a id="L87"></a>key, ok1 := Value(concrete.Key(), rand);
            <a id="L88"></a>value, ok2 := Value(concrete.Elem(), rand);
            <a id="L89"></a>if !ok1 || !ok2 {
                <a id="L90"></a>return nil, false
            <a id="L91"></a>}
            <a id="L92"></a>m.SetElem(key, value);
        <a id="L93"></a>}
        <a id="L94"></a>return m, true;
    <a id="L95"></a>case *reflect.PtrType:
        <a id="L96"></a>v, ok := Value(concrete.Elem(), rand);
        <a id="L97"></a>if !ok {
            <a id="L98"></a>return nil, false
        <a id="L99"></a>}
        <a id="L100"></a>p := reflect.MakeZero(concrete);
        <a id="L101"></a>p.(*reflect.PtrValue).PointTo(v);
        <a id="L102"></a>return p, true;
    <a id="L103"></a>case *reflect.SliceType:
        <a id="L104"></a>numElems := rand.Intn(complexSize);
        <a id="L105"></a>s := reflect.MakeSlice(concrete, numElems, numElems);
        <a id="L106"></a>for i := 0; i &lt; numElems; i++ {
            <a id="L107"></a>v, ok := Value(concrete.Elem(), rand);
            <a id="L108"></a>if !ok {
                <a id="L109"></a>return nil, false
            <a id="L110"></a>}
            <a id="L111"></a>s.Elem(i).SetValue(v);
        <a id="L112"></a>}
        <a id="L113"></a>return s, true;
    <a id="L114"></a>case *reflect.StringType:
        <a id="L115"></a>numChars := rand.Intn(complexSize);
        <a id="L116"></a>codePoints := make([]int, numChars);
        <a id="L117"></a>for i := 0; i &lt; numChars; i++ {
            <a id="L118"></a>codePoints[i] = rand.Intn(0x10ffff)
        <a id="L119"></a>}
        <a id="L120"></a>return reflect.NewValue(string(codePoints)), true;
    <a id="L121"></a>case *reflect.StructType:
        <a id="L122"></a>s := reflect.MakeZero(t).(*reflect.StructValue);
        <a id="L123"></a>for i := 0; i &lt; s.NumField(); i++ {
            <a id="L124"></a>v, ok := Value(concrete.Field(i).Type, rand);
            <a id="L125"></a>if !ok {
                <a id="L126"></a>return nil, false
            <a id="L127"></a>}
            <a id="L128"></a>s.Field(i).SetValue(v);
        <a id="L129"></a>}
        <a id="L130"></a>return s, true;
    <a id="L131"></a>case *reflect.Uint16Type:
        <a id="L132"></a>return reflect.NewValue(uint16(randInt64(rand))), true
    <a id="L133"></a>case *reflect.Uint32Type:
        <a id="L134"></a>return reflect.NewValue(uint32(randInt64(rand))), true
    <a id="L135"></a>case *reflect.Uint64Type:
        <a id="L136"></a>return reflect.NewValue(uint64(randInt64(rand))), true
    <a id="L137"></a>case *reflect.Uint8Type:
        <a id="L138"></a>return reflect.NewValue(uint8(randInt64(rand))), true
    <a id="L139"></a>case *reflect.UintType:
        <a id="L140"></a>return reflect.NewValue(uint(randInt64(rand))), true
    <a id="L141"></a>case *reflect.UintptrType:
        <a id="L142"></a>return reflect.NewValue(uintptr(randInt64(rand))), true
    <a id="L143"></a>default:
        <a id="L144"></a>return nil, false
    <a id="L145"></a>}

    <a id="L147"></a>return;
<a id="L148"></a>}

<a id="L150"></a><span class="comment">// A Config structure contains options for running a test.</span>
<a id="L151"></a>type Config struct {
    <a id="L152"></a><span class="comment">// MaxCount sets the maximum number of iterations. If zero,</span>
    <a id="L153"></a><span class="comment">// MaxCountScale is used.</span>
    <a id="L154"></a>MaxCount int;
    <a id="L155"></a><span class="comment">// MaxCountScale is a non-negative scale factor applied to the default</span>
    <a id="L156"></a><span class="comment">// maximum. If zero, the default is unchanged.</span>
    <a id="L157"></a>MaxCountScale float;
    <a id="L158"></a><span class="comment">// If non-nil, rand is a source of random numbers. Otherwise a default</span>
    <a id="L159"></a><span class="comment">// pseudo-random source will be used.</span>
    <a id="L160"></a>Rand *rand.Rand;
    <a id="L161"></a><span class="comment">// If non-nil, Values is a function which generates a slice of arbitrary</span>
    <a id="L162"></a><span class="comment">// Values that are congruent with the arguments to the function being</span>
    <a id="L163"></a><span class="comment">// tested. Otherwise, Values is used to generate the values.</span>
    <a id="L164"></a>Values func([]reflect.Value, *rand.Rand);
<a id="L165"></a>}

<a id="L167"></a>var defaultConfig Config

<a id="L169"></a><span class="comment">// getRand returns the *rand.Rand to use for a given Config.</span>
<a id="L170"></a>func (c *Config) getRand() *rand.Rand {
    <a id="L171"></a>if c.Rand == nil {
        <a id="L172"></a>return rand.New(rand.NewSource(0))
    <a id="L173"></a>}
    <a id="L174"></a>return c.Rand;
<a id="L175"></a>}

<a id="L177"></a><span class="comment">// getMaxCount returns the maximum number of iterations to run for a given</span>
<a id="L178"></a><span class="comment">// Config.</span>
<a id="L179"></a>func (c *Config) getMaxCount() (maxCount int) {
    <a id="L180"></a>maxCount = c.MaxCount;
    <a id="L181"></a>if maxCount == 0 {
        <a id="L182"></a>if c.MaxCountScale != 0 {
            <a id="L183"></a>maxCount = int(c.MaxCountScale * float(*defaultMaxCount))
        <a id="L184"></a>} else {
            <a id="L185"></a>maxCount = *defaultMaxCount
        <a id="L186"></a>}
    <a id="L187"></a>}

    <a id="L189"></a>return;
<a id="L190"></a>}

<a id="L192"></a><span class="comment">// A SetupError is the result of an error in the way that check is being</span>
<a id="L193"></a><span class="comment">// used, independent of the functions being tested.</span>
<a id="L194"></a>type SetupError string

<a id="L196"></a>func (s SetupError) String() string { return string(s) }

<a id="L198"></a><span class="comment">// A CheckError is the result of Check finding an error.</span>
<a id="L199"></a>type CheckError struct {
    <a id="L200"></a>Count int;
    <a id="L201"></a>In    []interface{};
<a id="L202"></a>}

<a id="L204"></a>func (s *CheckError) String() string {
    <a id="L205"></a>return fmt.Sprintf(&#34;#%d: failed on input %s&#34;, s.Count, toString(s.In))
<a id="L206"></a>}

<a id="L208"></a><span class="comment">// A CheckEqualError is the result CheckEqual finding an error.</span>
<a id="L209"></a>type CheckEqualError struct {
    <a id="L210"></a>CheckError;
    <a id="L211"></a>Out1 []interface{};
    <a id="L212"></a>Out2 []interface{};
<a id="L213"></a>}

<a id="L215"></a>func (s *CheckEqualError) String() string {
    <a id="L216"></a>return fmt.Sprintf(&#34;#%d: failed on input %s. Output 1: %s. Output 2: %s&#34;, s.Count, toString(s.In), toString(s.Out1), toString(s.Out2))
<a id="L217"></a>}

<a id="L219"></a><span class="comment">// Check looks for an input to f, any function that returns bool,</span>
<a id="L220"></a><span class="comment">// such that f returns false.  It calls f repeatedly, with arbitrary</span>
<a id="L221"></a><span class="comment">// values for each argument.  If f returns false on a given input,</span>
<a id="L222"></a><span class="comment">// Check returns that input as a *CheckError.</span>
<a id="L223"></a><span class="comment">// For example:</span>
<a id="L224"></a><span class="comment">//</span>
<a id="L225"></a><span class="comment">// 	func TestOddMultipleOfThree(t *testing.T) {</span>
<a id="L226"></a><span class="comment">// 		f := func(x int) bool {</span>
<a id="L227"></a><span class="comment">// 			y := OddMultipleOfThree(x);</span>
<a id="L228"></a><span class="comment">// 			return y%2 == 1 &amp;&amp; y%3 == 0</span>
<a id="L229"></a><span class="comment">// 		}</span>
<a id="L230"></a><span class="comment">// 		if err := quick.Check(f, nil); err != nil {</span>
<a id="L231"></a><span class="comment">// 			t.Error(err);</span>
<a id="L232"></a><span class="comment">// 		}</span>
<a id="L233"></a><span class="comment">// 	}</span>
<a id="L234"></a>func Check(function interface{}, config *Config) (err os.Error) {
    <a id="L235"></a>if config == nil {
        <a id="L236"></a>config = &amp;defaultConfig
    <a id="L237"></a>}

    <a id="L239"></a>f, fType, ok := functionAndType(function);
    <a id="L240"></a>if !ok {
        <a id="L241"></a>err = SetupError(&#34;argument is not a function&#34;);
        <a id="L242"></a>return;
    <a id="L243"></a>}

    <a id="L245"></a>if fType.NumOut() != 1 {
        <a id="L246"></a>err = SetupError(&#34;function returns more than one value.&#34;);
        <a id="L247"></a>return;
    <a id="L248"></a>}
    <a id="L249"></a>if _, ok := fType.Out(0).(*reflect.BoolType); !ok {
        <a id="L250"></a>err = SetupError(&#34;function does not return a bool&#34;);
        <a id="L251"></a>return;
    <a id="L252"></a>}

    <a id="L254"></a>arguments := make([]reflect.Value, fType.NumIn());
    <a id="L255"></a>rand := config.getRand();
    <a id="L256"></a>maxCount := config.getMaxCount();

    <a id="L258"></a>for i := 0; i &lt; maxCount; i++ {
        <a id="L259"></a>err = arbitraryValues(arguments, fType, config, rand);
        <a id="L260"></a>if err != nil {
            <a id="L261"></a>return
        <a id="L262"></a>}

        <a id="L264"></a>if !f.Call(arguments)[0].(*reflect.BoolValue).Get() {
            <a id="L265"></a>err = &amp;CheckError{i + 1, toInterfaces(arguments)};
            <a id="L266"></a>return;
        <a id="L267"></a>}
    <a id="L268"></a>}

    <a id="L270"></a>return;
<a id="L271"></a>}

<a id="L273"></a><span class="comment">// CheckEqual looks for an input on which f and g return different results.</span>
<a id="L274"></a><span class="comment">// It calls f and g repeatedly with arbitrary values for each argument.</span>
<a id="L275"></a><span class="comment">// If f and g return different answers, CheckEqual returns a *CheckEqualError</span>
<a id="L276"></a><span class="comment">// describing the input and the outputs.</span>
<a id="L277"></a>func CheckEqual(f, g interface{}, config *Config) (err os.Error) {
    <a id="L278"></a>if config == nil {
        <a id="L279"></a>config = &amp;defaultConfig
    <a id="L280"></a>}

    <a id="L282"></a>x, xType, ok := functionAndType(f);
    <a id="L283"></a>if !ok {
        <a id="L284"></a>err = SetupError(&#34;f is not a function&#34;);
        <a id="L285"></a>return;
    <a id="L286"></a>}
    <a id="L287"></a>y, yType, ok := functionAndType(g);
    <a id="L288"></a>if !ok {
        <a id="L289"></a>err = SetupError(&#34;g is not a function&#34;);
        <a id="L290"></a>return;
    <a id="L291"></a>}

    <a id="L293"></a>if xType != yType {
        <a id="L294"></a>err = SetupError(&#34;functions have different types&#34;);
        <a id="L295"></a>return;
    <a id="L296"></a>}

    <a id="L298"></a>arguments := make([]reflect.Value, xType.NumIn());
    <a id="L299"></a>rand := config.getRand();
    <a id="L300"></a>maxCount := config.getMaxCount();

    <a id="L302"></a>for i := 0; i &lt; maxCount; i++ {
        <a id="L303"></a>err = arbitraryValues(arguments, xType, config, rand);
        <a id="L304"></a>if err != nil {
            <a id="L305"></a>return
        <a id="L306"></a>}

        <a id="L308"></a>xOut := toInterfaces(x.Call(arguments));
        <a id="L309"></a>yOut := toInterfaces(y.Call(arguments));

        <a id="L311"></a>if !reflect.DeepEqual(xOut, yOut) {
            <a id="L312"></a>err = &amp;CheckEqualError{CheckError{i + 1, toInterfaces(arguments)}, xOut, yOut};
            <a id="L313"></a>return;
        <a id="L314"></a>}
    <a id="L315"></a>}

    <a id="L317"></a>return;
<a id="L318"></a>}

<a id="L320"></a><span class="comment">// arbitraryValues writes Values to args such that args contains Values</span>
<a id="L321"></a><span class="comment">// suitable for calling f.</span>
<a id="L322"></a>func arbitraryValues(args []reflect.Value, f *reflect.FuncType, config *Config, rand *rand.Rand) (err os.Error) {
    <a id="L323"></a>if config.Values != nil {
        <a id="L324"></a>config.Values(args, rand);
        <a id="L325"></a>return;
    <a id="L326"></a>}

    <a id="L328"></a>for j := 0; j &lt; len(args); j++ {
        <a id="L329"></a>var ok bool;
        <a id="L330"></a>args[j], ok = Value(f.In(j), rand);
        <a id="L331"></a>if !ok {
            <a id="L332"></a>err = SetupError(fmt.Sprintf(&#34;cannot create arbitrary value of type %s for argument %d&#34;, f.In(j), j));
            <a id="L333"></a>return;
        <a id="L334"></a>}
    <a id="L335"></a>}

    <a id="L337"></a>return;
<a id="L338"></a>}

<a id="L340"></a>func functionAndType(f interface{}) (v *reflect.FuncValue, t *reflect.FuncType, ok bool) {
    <a id="L341"></a>v, ok = reflect.NewValue(f).(*reflect.FuncValue);
    <a id="L342"></a>if !ok {
        <a id="L343"></a>return
    <a id="L344"></a>}
    <a id="L345"></a>t = v.Type().(*reflect.FuncType);
    <a id="L346"></a>return;
<a id="L347"></a>}

<a id="L349"></a>func toInterfaces(values []reflect.Value) []interface{} {
    <a id="L350"></a>ret := make([]interface{}, len(values));
    <a id="L351"></a>for i, v := range values {
        <a id="L352"></a>ret[i] = v.Interface()
    <a id="L353"></a>}
    <a id="L354"></a>return ret;
<a id="L355"></a>}

<a id="L357"></a>func toString(interfaces []interface{}) string {
    <a id="L358"></a>s := make([]string, len(interfaces));
    <a id="L359"></a>for i, v := range interfaces {
        <a id="L360"></a>s[i] = fmt.Sprintf(&#34;%#v&#34;, v)
    <a id="L361"></a>}
    <a id="L362"></a>return strings.Join(s, &#34;, &#34;);
<a id="L363"></a>}
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
