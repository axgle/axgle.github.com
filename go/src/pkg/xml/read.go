<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/xml/read.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/xml/read.go</h1>

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
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;io&#34;;
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;reflect&#34;;
    <a id="L12"></a>&#34;strings&#34;;
<a id="L13"></a>)

<a id="L15"></a><span class="comment">// BUG(rsc): Mapping between XML elements and data structures is inherently flawed:</span>
<a id="L16"></a><span class="comment">// an XML element is an order-dependent collection of anonymous</span>
<a id="L17"></a><span class="comment">// values, while a data structure is an order-independent collection</span>
<a id="L18"></a><span class="comment">// of named values.</span>
<a id="L19"></a><span class="comment">// See package json for a textual representation more suitable</span>
<a id="L20"></a><span class="comment">// to data structures.</span>

<a id="L22"></a><span class="comment">// Unmarshal parses an XML element from r and uses the</span>
<a id="L23"></a><span class="comment">// reflect library to fill in an arbitrary struct, slice, or string</span>
<a id="L24"></a><span class="comment">// pointed at by val.  Well-formed data that does not fit</span>
<a id="L25"></a><span class="comment">// into val is discarded.</span>
<a id="L26"></a><span class="comment">//</span>
<a id="L27"></a><span class="comment">// For example, given these definitions:</span>
<a id="L28"></a><span class="comment">//</span>
<a id="L29"></a><span class="comment">//	type Email struct {</span>
<a id="L30"></a><span class="comment">//		Where string &#34;attr&#34;;</span>
<a id="L31"></a><span class="comment">//		Addr string;</span>
<a id="L32"></a><span class="comment">//	}</span>
<a id="L33"></a><span class="comment">//</span>
<a id="L34"></a><span class="comment">//	type Result struct {</span>
<a id="L35"></a><span class="comment">//		XMLName xml.Name &#34;result&#34;;</span>
<a id="L36"></a><span class="comment">//		Name string;</span>
<a id="L37"></a><span class="comment">//		Phone string;</span>
<a id="L38"></a><span class="comment">//		Email []Email;</span>
<a id="L39"></a><span class="comment">//	}</span>
<a id="L40"></a><span class="comment">//</span>
<a id="L41"></a><span class="comment">//	var result = Result{ &#34;name&#34;, &#34;phone&#34;, nil }</span>
<a id="L42"></a><span class="comment">//</span>
<a id="L43"></a><span class="comment">// unmarshalling the XML input</span>
<a id="L44"></a><span class="comment">//</span>
<a id="L45"></a><span class="comment">//	&lt;result&gt;</span>
<a id="L46"></a><span class="comment">//		&lt;email where=&#34;home&#34;&gt;</span>
<a id="L47"></a><span class="comment">//			&lt;addr&gt;gre@example.com&lt;/addr&gt;</span>
<a id="L48"></a><span class="comment">//		&lt;/email&gt;</span>
<a id="L49"></a><span class="comment">//		&lt;email where=&#39;work&#39;&gt;</span>
<a id="L50"></a><span class="comment">//			&lt;addr&gt;gre@work.com&lt;/addr&gt;</span>
<a id="L51"></a><span class="comment">//		&lt;/email&gt;</span>
<a id="L52"></a><span class="comment">//		&lt;name&gt;Grace R. Emlin&lt;/name&gt;</span>
<a id="L53"></a><span class="comment">//		&lt;address&gt;123 Main Street&lt;/address&gt;</span>
<a id="L54"></a><span class="comment">//	&lt;/result&gt;</span>
<a id="L55"></a><span class="comment">//</span>
<a id="L56"></a><span class="comment">// via Unmarshal(r, &amp;result) is equivalent to assigning</span>
<a id="L57"></a><span class="comment">//</span>
<a id="L58"></a><span class="comment">//	r = Result{</span>
<a id="L59"></a><span class="comment">//		xml.Name{&#34;&#34;, &#34;result&#34;},</span>
<a id="L60"></a><span class="comment">//		&#34;Grace R. Emlin&#34;,	// name</span>
<a id="L61"></a><span class="comment">//		&#34;phone&#34;,	// no phone given</span>
<a id="L62"></a><span class="comment">//		[]Email{</span>
<a id="L63"></a><span class="comment">//			Email{ &#34;home&#34;, &#34;gre@example.com&#34; },</span>
<a id="L64"></a><span class="comment">//			Email{ &#34;work&#34;, &#34;gre@work.com&#34; }</span>
<a id="L65"></a><span class="comment">//		}</span>
<a id="L66"></a><span class="comment">//	}</span>
<a id="L67"></a><span class="comment">//</span>
<a id="L68"></a><span class="comment">// Note that the field r.Phone has not been modified and</span>
<a id="L69"></a><span class="comment">// that the XML &lt;address&gt; element was discarded.</span>
<a id="L70"></a><span class="comment">//</span>
<a id="L71"></a><span class="comment">// Because Unmarshal uses the reflect package, it can only</span>
<a id="L72"></a><span class="comment">// assign to upper case fields.  Unmarshal uses a case-insensitive</span>
<a id="L73"></a><span class="comment">// comparison to match XML element names to struct field names.</span>
<a id="L74"></a><span class="comment">//</span>
<a id="L75"></a><span class="comment">// Unmarshal maps an XML element to a struct using the following rules:</span>
<a id="L76"></a><span class="comment">//</span>
<a id="L77"></a><span class="comment">//   * If the struct has a field named XMLName of type xml.Name,</span>
<a id="L78"></a><span class="comment">//      Unmarshal records the element name in that field.</span>
<a id="L79"></a><span class="comment">//</span>
<a id="L80"></a><span class="comment">//   * If the XMLName field has an associated tag string of the form</span>
<a id="L81"></a><span class="comment">//      &#34;tag&#34; or &#34;namespace-URL tag&#34;, the XML element must have</span>
<a id="L82"></a><span class="comment">//      the given tag (and, optionally, name space) or else Unmarshal</span>
<a id="L83"></a><span class="comment">//      returns an error.</span>
<a id="L84"></a><span class="comment">//</span>
<a id="L85"></a><span class="comment">//   * If the XML element has an attribute whose name matches a</span>
<a id="L86"></a><span class="comment">//      struct field of type string with tag &#34;attr&#34;, Unmarshal records</span>
<a id="L87"></a><span class="comment">//      the attribute value in that field.</span>
<a id="L88"></a><span class="comment">//</span>
<a id="L89"></a><span class="comment">//   * If the XML element contains character data, that data is</span>
<a id="L90"></a><span class="comment">//      accumulated in the first struct field that has tag &#34;chardata&#34;.</span>
<a id="L91"></a><span class="comment">//      The struct field may have type []byte or string.</span>
<a id="L92"></a><span class="comment">//      If there is no such field, the character data is discarded.</span>
<a id="L93"></a><span class="comment">//</span>
<a id="L94"></a><span class="comment">//   * If the XML element contains a sub-element whose name</span>
<a id="L95"></a><span class="comment">//      matches a struct field whose tag is neither &#34;attr&#34; nor &#34;chardata&#34;,</span>
<a id="L96"></a><span class="comment">//      Unmarshal maps the sub-element to that struct field.</span>
<a id="L97"></a><span class="comment">//      Otherwise, if the struct has a field named Any, unmarshal</span>
<a id="L98"></a><span class="comment">//      maps the sub-element to that struct field.</span>
<a id="L99"></a><span class="comment">//</span>
<a id="L100"></a><span class="comment">// Unmarshal maps an XML element to a string or []byte by saving the</span>
<a id="L101"></a><span class="comment">// concatenation of that elements character data in the string or []byte.</span>
<a id="L102"></a><span class="comment">//</span>
<a id="L103"></a><span class="comment">// Unmarshal maps an XML element to a slice by extending the length</span>
<a id="L104"></a><span class="comment">// of the slice and mapping the element to the newly created value.</span>
<a id="L105"></a><span class="comment">//</span>
<a id="L106"></a><span class="comment">// Unmarshal maps an XML element to a bool by setting the bool to true.</span>
<a id="L107"></a><span class="comment">//</span>
<a id="L108"></a><span class="comment">// Unmarshal maps an XML element to an xml.Name by recording the</span>
<a id="L109"></a><span class="comment">// element name.</span>
<a id="L110"></a><span class="comment">//</span>
<a id="L111"></a><span class="comment">// Unmarshal maps an XML element to a pointer by setting the pointer</span>
<a id="L112"></a><span class="comment">// to a freshly allocated value and then mapping the element to that value.</span>
<a id="L113"></a><span class="comment">//</span>
<a id="L114"></a>func Unmarshal(r io.Reader, val interface{}) os.Error {
    <a id="L115"></a>v, ok := reflect.NewValue(val).(*reflect.PtrValue);
    <a id="L116"></a>if !ok {
        <a id="L117"></a>return os.NewError(&#34;non-pointer passed to Unmarshal&#34;)
    <a id="L118"></a>}
    <a id="L119"></a>p := NewParser(r);
    <a id="L120"></a>elem := v.Elem();
    <a id="L121"></a>err := p.unmarshal(elem, nil);
    <a id="L122"></a>if err != nil {
        <a id="L123"></a>return err
    <a id="L124"></a>}
    <a id="L125"></a>return nil;
<a id="L126"></a>}

<a id="L128"></a><span class="comment">// An UnmarshalError represents an error in the unmarshalling process.</span>
<a id="L129"></a>type UnmarshalError string

<a id="L131"></a>func (e UnmarshalError) String() string { return string(e) }

<a id="L133"></a><span class="comment">// The Parser&#39;s Unmarshal method is like xml.Unmarshal</span>
<a id="L134"></a><span class="comment">// except that it can be passed a pointer to the initial start element,</span>
<a id="L135"></a><span class="comment">// useful when a client reads some raw XML tokens itself</span>
<a id="L136"></a><span class="comment">// but also defers to Unmarshal for some elements.</span>
<a id="L137"></a><span class="comment">// Passing a nil start element indicates that Unmarshal should</span>
<a id="L138"></a><span class="comment">// read the token stream to find the start element.</span>
<a id="L139"></a>func (p *Parser) Unmarshal(val interface{}, start *StartElement) os.Error {
    <a id="L140"></a>v, ok := reflect.NewValue(val).(*reflect.PtrValue);
    <a id="L141"></a>if !ok {
        <a id="L142"></a>return os.NewError(&#34;non-pointer passed to Unmarshal&#34;)
    <a id="L143"></a>}
    <a id="L144"></a>return p.unmarshal(v.Elem(), start);
<a id="L145"></a>}

<a id="L147"></a><span class="comment">// Unmarshal a single XML element into val.</span>
<a id="L148"></a>func (p *Parser) unmarshal(val reflect.Value, start *StartElement) os.Error {
    <a id="L149"></a><span class="comment">// Find start element if we need it.</span>
    <a id="L150"></a>if start == nil {
        <a id="L151"></a>for {
            <a id="L152"></a>tok, err := p.Token();
            <a id="L153"></a>if err != nil {
                <a id="L154"></a>return err
            <a id="L155"></a>}
            <a id="L156"></a>if t, ok := tok.(StartElement); ok {
                <a id="L157"></a>start = &amp;t;
                <a id="L158"></a>break;
            <a id="L159"></a>}
        <a id="L160"></a>}
    <a id="L161"></a>}

    <a id="L163"></a>if pv, ok := val.(*reflect.PtrValue); ok {
        <a id="L164"></a>if pv.Get() == 0 {
            <a id="L165"></a>zv := reflect.MakeZero(pv.Type().(*reflect.PtrType).Elem());
            <a id="L166"></a>pv.PointTo(zv);
            <a id="L167"></a>val = zv;
        <a id="L168"></a>} else {
            <a id="L169"></a>val = pv.Elem()
        <a id="L170"></a>}
    <a id="L171"></a>}

    <a id="L173"></a>var (
        <a id="L174"></a>data        []byte;
        <a id="L175"></a>saveData    reflect.Value;
        <a id="L176"></a>comment     []byte;
        <a id="L177"></a>saveComment reflect.Value;
        <a id="L178"></a>sv          *reflect.StructValue;
        <a id="L179"></a>styp        *reflect.StructType;
    <a id="L180"></a>)
    <a id="L181"></a>switch v := val.(type) {
    <a id="L182"></a>case *reflect.BoolValue:
        <a id="L183"></a>v.Set(true)

    <a id="L185"></a>case *reflect.SliceValue:
        <a id="L186"></a>typ := v.Type().(*reflect.SliceType);
        <a id="L187"></a>if _, ok := typ.Elem().(*reflect.Uint8Type); ok {
            <a id="L188"></a><span class="comment">// []byte</span>
            <a id="L189"></a>saveData = v;
            <a id="L190"></a>break;
        <a id="L191"></a>}

        <a id="L193"></a><span class="comment">// Slice of element values.</span>
        <a id="L194"></a><span class="comment">// Grow slice.</span>
        <a id="L195"></a>n := v.Len();
        <a id="L196"></a>if n &gt;= v.Cap() {
            <a id="L197"></a>ncap := 2 * n;
            <a id="L198"></a>if ncap &lt; 4 {
                <a id="L199"></a>ncap = 4
            <a id="L200"></a>}
            <a id="L201"></a>new := reflect.MakeSlice(typ, n, ncap);
            <a id="L202"></a>reflect.ArrayCopy(new, v);
            <a id="L203"></a>v.Set(new);
        <a id="L204"></a>}
        <a id="L205"></a>v.SetLen(n + 1);

        <a id="L207"></a><span class="comment">// Recur to read element into slice.</span>
        <a id="L208"></a>if err := p.unmarshal(v.Elem(n), start); err != nil {
            <a id="L209"></a>v.SetLen(n);
            <a id="L210"></a>return err;
        <a id="L211"></a>}
        <a id="L212"></a>return nil;

    <a id="L214"></a>case *reflect.StringValue:
        <a id="L215"></a>saveData = v

    <a id="L217"></a>case *reflect.StructValue:
        <a id="L218"></a>if _, ok := v.Interface().(Name); ok {
            <a id="L219"></a>v.Set(reflect.NewValue(start.Name).(*reflect.StructValue));
            <a id="L220"></a>break;
        <a id="L221"></a>}

        <a id="L223"></a>sv = v;
        <a id="L224"></a>typ := sv.Type().(*reflect.StructType);
        <a id="L225"></a>styp = typ;
        <a id="L226"></a><span class="comment">// Assign name.</span>
        <a id="L227"></a>if f, ok := typ.FieldByName(&#34;XMLName&#34;); ok {
            <a id="L228"></a><span class="comment">// Validate element name.</span>
            <a id="L229"></a>if f.Tag != &#34;&#34; {
                <a id="L230"></a>tag := f.Tag;
                <a id="L231"></a>ns := &#34;&#34;;
                <a id="L232"></a>i := strings.LastIndex(tag, &#34; &#34;);
                <a id="L233"></a>if i &gt;= 0 {
                    <a id="L234"></a>ns, tag = tag[0:i], tag[i+1:len(tag)]
                <a id="L235"></a>}
                <a id="L236"></a>if tag != start.Name.Local {
                    <a id="L237"></a>return UnmarshalError(&#34;expected element type &lt;&#34; + tag + &#34;&gt; but have &lt;&#34; + start.Name.Local + &#34;&gt;&#34;)
                <a id="L238"></a>}
                <a id="L239"></a>if ns != &#34;&#34; &amp;&amp; ns != start.Name.Space {
                    <a id="L240"></a>e := &#34;expected element &lt;&#34; + tag + &#34;&gt; in name space &#34; + ns + &#34; but have &#34;;
                    <a id="L241"></a>if start.Name.Space == &#34;&#34; {
                        <a id="L242"></a>e += &#34;no name space&#34;
                    <a id="L243"></a>} else {
                        <a id="L244"></a>e += start.Name.Space
                    <a id="L245"></a>}
                    <a id="L246"></a>return UnmarshalError(e);
                <a id="L247"></a>}
            <a id="L248"></a>}

            <a id="L250"></a><span class="comment">// Save</span>
            <a id="L251"></a>v := sv.FieldByIndex(f.Index);
            <a id="L252"></a>if _, ok := v.Interface().(Name); !ok {
                <a id="L253"></a>return UnmarshalError(sv.Type().String() + &#34; field XMLName does not have type xml.Name&#34;)
            <a id="L254"></a>}
            <a id="L255"></a>v.(*reflect.StructValue).Set(reflect.NewValue(start.Name).(*reflect.StructValue));
        <a id="L256"></a>}

        <a id="L258"></a><span class="comment">// Assign attributes.</span>
        <a id="L259"></a><span class="comment">// Also, determine whether we need to save character data or comments.</span>
        <a id="L260"></a>for i, n := 0, typ.NumField(); i &lt; n; i++ {
            <a id="L261"></a>f := typ.Field(i);
            <a id="L262"></a>switch f.Tag {
            <a id="L263"></a>case &#34;attr&#34;:
                <a id="L264"></a>strv, ok := sv.FieldByIndex(f.Index).(*reflect.StringValue);
                <a id="L265"></a>if !ok {
                    <a id="L266"></a>return UnmarshalError(sv.Type().String() + &#34; field &#34; + f.Name + &#34; has attr tag but is not type string&#34;)
                <a id="L267"></a>}
                <a id="L268"></a><span class="comment">// Look for attribute.</span>
                <a id="L269"></a>val := &#34;&#34;;
                <a id="L270"></a>k := strings.ToLower(f.Name);
                <a id="L271"></a>for _, a := range start.Attr {
                    <a id="L272"></a>if strings.ToLower(a.Name.Local) == k {
                        <a id="L273"></a>val = a.Value;
                        <a id="L274"></a>break;
                    <a id="L275"></a>}
                <a id="L276"></a>}
                <a id="L277"></a>strv.Set(val);

            <a id="L279"></a>case &#34;comment&#34;:
                <a id="L280"></a>if saveComment == nil {
                    <a id="L281"></a>saveComment = sv.FieldByIndex(f.Index)
                <a id="L282"></a>}

            <a id="L284"></a>case &#34;chardata&#34;:
                <a id="L285"></a>if saveData == nil {
                    <a id="L286"></a>saveData = sv.FieldByIndex(f.Index)
                <a id="L287"></a>}
            <a id="L288"></a>}
        <a id="L289"></a>}
    <a id="L290"></a>}

    <a id="L292"></a><span class="comment">// Find end element.</span>
    <a id="L293"></a><span class="comment">// Process sub-elements along the way.</span>
<a id="L294"></a>Loop:
    <a id="L295"></a>for {
        <a id="L296"></a>tok, err := p.Token();
        <a id="L297"></a>if err != nil {
            <a id="L298"></a>return err
        <a id="L299"></a>}
        <a id="L300"></a>switch t := tok.(type) {
        <a id="L301"></a>case StartElement:
            <a id="L302"></a><span class="comment">// Sub-element.</span>
            <a id="L303"></a><span class="comment">// Look up by tag name.</span>
            <a id="L304"></a><span class="comment">// If that fails, fall back to mop-up field named &#34;Any&#34;.</span>
            <a id="L305"></a>if sv != nil {
                <a id="L306"></a>k := strings.ToLower(t.Name.Local);
                <a id="L307"></a>any := -1;
                <a id="L308"></a>for i, n := 0, styp.NumField(); i &lt; n; i++ {
                    <a id="L309"></a>f := styp.Field(i);
                    <a id="L310"></a>if strings.ToLower(f.Name) == k {
                        <a id="L311"></a>if err := p.unmarshal(sv.FieldByIndex(f.Index), &amp;t); err != nil {
                            <a id="L312"></a>return err
                        <a id="L313"></a>}
                        <a id="L314"></a>continue Loop;
                    <a id="L315"></a>}
                    <a id="L316"></a>if any &lt; 0 &amp;&amp; f.Name == &#34;Any&#34; {
                        <a id="L317"></a>any = i
                    <a id="L318"></a>}
                <a id="L319"></a>}
                <a id="L320"></a>if any &gt;= 0 {
                    <a id="L321"></a>if err := p.unmarshal(sv.FieldByIndex(styp.Field(any).Index), &amp;t); err != nil {
                        <a id="L322"></a>return err
                    <a id="L323"></a>}
                    <a id="L324"></a>continue Loop;
                <a id="L325"></a>}
            <a id="L326"></a>}
            <a id="L327"></a><span class="comment">// Not saving sub-element but still have to skip over it.</span>
            <a id="L328"></a>if err := p.Skip(); err != nil {
                <a id="L329"></a>return err
            <a id="L330"></a>}

        <a id="L332"></a>case EndElement:
            <a id="L333"></a>break Loop

        <a id="L335"></a>case CharData:
            <a id="L336"></a>if saveData != nil {
                <a id="L337"></a>data = bytes.Add(data, t)
            <a id="L338"></a>}

        <a id="L340"></a>case Comment:
            <a id="L341"></a>if saveComment != nil {
                <a id="L342"></a>comment = bytes.Add(comment, t)
            <a id="L343"></a>}
        <a id="L344"></a>}
    <a id="L345"></a>}

    <a id="L347"></a><span class="comment">// Save accumulated character data and comments</span>
    <a id="L348"></a>switch t := saveData.(type) {
    <a id="L349"></a>case *reflect.StringValue:
        <a id="L350"></a>t.Set(string(data))
    <a id="L351"></a>case *reflect.SliceValue:
        <a id="L352"></a>t.Set(reflect.NewValue(data).(*reflect.SliceValue))
    <a id="L353"></a>}

    <a id="L355"></a>switch t := saveComment.(type) {
    <a id="L356"></a>case *reflect.StringValue:
        <a id="L357"></a>t.Set(string(comment))
    <a id="L358"></a>case *reflect.SliceValue:
        <a id="L359"></a>t.Set(reflect.NewValue(comment).(*reflect.SliceValue))
    <a id="L360"></a>}

    <a id="L362"></a>return nil;
<a id="L363"></a>}

<a id="L365"></a><span class="comment">// Have already read a start element.</span>
<a id="L366"></a><span class="comment">// Read tokens until we find the end element.</span>
<a id="L367"></a><span class="comment">// Token is taking care of making sure the</span>
<a id="L368"></a><span class="comment">// end element matches the start element we saw.</span>
<a id="L369"></a>func (p *Parser) Skip() os.Error {
    <a id="L370"></a>for {
        <a id="L371"></a>tok, err := p.Token();
        <a id="L372"></a>if err != nil {
            <a id="L373"></a>return err
        <a id="L374"></a>}
        <a id="L375"></a>switch t := tok.(type) {
        <a id="L376"></a>case StartElement:
            <a id="L377"></a>if err := p.Skip(); err != nil {
                <a id="L378"></a>return err
            <a id="L379"></a>}
        <a id="L380"></a>case EndElement:
            <a id="L381"></a>return nil
        <a id="L382"></a>}
    <a id="L383"></a>}
    <a id="L384"></a>panic(&#34;unreachable&#34;);
<a id="L385"></a>}
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
