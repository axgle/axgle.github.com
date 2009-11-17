<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/rsa/rsa.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/rsa/rsa.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package implements RSA encryption as specified in PKCS#1.</span>
<a id="L6"></a>package rsa

<a id="L8"></a><span class="comment">// TODO(agl): Add support for PSS padding.</span>

<a id="L10"></a>import (
    <a id="L11"></a>&#34;big&#34;;
    <a id="L12"></a>&#34;bytes&#34;;
    <a id="L13"></a>&#34;crypto/subtle&#34;;
    <a id="L14"></a>&#34;hash&#34;;
    <a id="L15"></a>&#34;io&#34;;
    <a id="L16"></a>&#34;os&#34;;
<a id="L17"></a>)

<a id="L19"></a>var bigZero = big.NewInt(0)
<a id="L20"></a>var bigOne = big.NewInt(1)

<a id="L22"></a><span class="comment">// randomSafePrime returns a number, p, of the given size, such that p and</span>
<a id="L23"></a><span class="comment">// (p-1)/2 are both prime with high probability.</span>
<a id="L24"></a>func randomSafePrime(rand io.Reader, bits int) (p *big.Int, err os.Error) {
    <a id="L25"></a>if bits &lt; 1 {
        <a id="L26"></a>err = os.EINVAL
    <a id="L27"></a>}

    <a id="L29"></a>bytes := make([]byte, (bits+7)/8);
    <a id="L30"></a>p = new(big.Int);
    <a id="L31"></a>p2 := new(big.Int);

    <a id="L33"></a>for {
        <a id="L34"></a>_, err = io.ReadFull(rand, bytes);
        <a id="L35"></a>if err != nil {
            <a id="L36"></a>return
        <a id="L37"></a>}

        <a id="L39"></a><span class="comment">// Don&#39;t let the value be too small.</span>
        <a id="L40"></a>bytes[0] |= 0x80;
        <a id="L41"></a><span class="comment">// Make the value odd since an even number this large certainly isn&#39;t prime.</span>
        <a id="L42"></a>bytes[len(bytes)-1] |= 1;

        <a id="L44"></a>p.SetBytes(bytes);
        <a id="L45"></a>if big.ProbablyPrime(p, 20) {
            <a id="L46"></a>p2.Rsh(p, 1); <span class="comment">// p2 = (p - 1)/2</span>
            <a id="L47"></a>if big.ProbablyPrime(p2, 20) {
                <a id="L48"></a>return
            <a id="L49"></a>}
        <a id="L50"></a>}
    <a id="L51"></a>}

    <a id="L53"></a>return;
<a id="L54"></a>}

<a id="L56"></a><span class="comment">// randomNumber returns a uniform random value in [0, max).</span>
<a id="L57"></a>func randomNumber(rand io.Reader, max *big.Int) (n *big.Int, err os.Error) {
    <a id="L58"></a>k := (max.Len() + 7) / 8;

    <a id="L60"></a><span class="comment">// r is the number of bits in the used in the most significant byte of</span>
    <a id="L61"></a><span class="comment">// max.</span>
    <a id="L62"></a>r := uint(max.Len() % 8);
    <a id="L63"></a>if r == 0 {
        <a id="L64"></a>r = 8
    <a id="L65"></a>}

    <a id="L67"></a>bytes := make([]byte, k);
    <a id="L68"></a>n = new(big.Int);

    <a id="L70"></a>for {
        <a id="L71"></a>_, err = io.ReadFull(rand, bytes);
        <a id="L72"></a>if err != nil {
            <a id="L73"></a>return
        <a id="L74"></a>}

        <a id="L76"></a><span class="comment">// Clear bits in the first byte to increase the probability</span>
        <a id="L77"></a><span class="comment">// that the candidate is &lt; max.</span>
        <a id="L78"></a>bytes[0] &amp;= uint8(int(1&lt;&lt;r) - 1);

        <a id="L80"></a>n.SetBytes(bytes);
        <a id="L81"></a>if n.Cmp(max) &lt; 0 {
            <a id="L82"></a>return
        <a id="L83"></a>}
    <a id="L84"></a>}

    <a id="L86"></a>return;
<a id="L87"></a>}

<a id="L89"></a><span class="comment">// A PublicKey represents the public part of an RSA key.</span>
<a id="L90"></a>type PublicKey struct {
    <a id="L91"></a>N   *big.Int; <span class="comment">// modulus</span>
    <a id="L92"></a>E   int;      <span class="comment">// public exponent</span>
<a id="L93"></a>}

<a id="L95"></a><span class="comment">// A PrivateKey represents an RSA key</span>
<a id="L96"></a>type PrivateKey struct {
    <a id="L97"></a>PublicKey;           <span class="comment">// public part.</span>
    <a id="L98"></a>D          *big.Int; <span class="comment">// private exponent</span>
    <a id="L99"></a>P, Q       *big.Int; <span class="comment">// prime factors of N</span>
<a id="L100"></a>}

<a id="L102"></a><span class="comment">// Validate performs basic sanity checks on the key.</span>
<a id="L103"></a><span class="comment">// It returns nil if the key is valid, or else an os.Error describing a problem.</span>

<a id="L105"></a>func (priv PrivateKey) Validate() os.Error {
    <a id="L106"></a><span class="comment">// Check that p and q are prime. Note that this is just a sanity</span>
    <a id="L107"></a><span class="comment">// check. Since the random witnesses chosen by ProbablyPrime are</span>
    <a id="L108"></a><span class="comment">// deterministic, given the candidate number, it&#39;s easy for an attack</span>
    <a id="L109"></a><span class="comment">// to generate composites that pass this test.</span>
    <a id="L110"></a>if !big.ProbablyPrime(priv.P, 20) {
        <a id="L111"></a>return os.ErrorString(&#34;P is composite&#34;)
    <a id="L112"></a>}
    <a id="L113"></a>if !big.ProbablyPrime(priv.Q, 20) {
        <a id="L114"></a>return os.ErrorString(&#34;Q is composite&#34;)
    <a id="L115"></a>}

    <a id="L117"></a><span class="comment">// Check that p*q == n.</span>
    <a id="L118"></a>modulus := new(big.Int).Mul(priv.P, priv.Q);
    <a id="L119"></a>if modulus.Cmp(priv.N) != 0 {
        <a id="L120"></a>return os.ErrorString(&#34;invalid modulus&#34;)
    <a id="L121"></a>}
    <a id="L122"></a><span class="comment">// Check that e and totient(p, q) are coprime.</span>
    <a id="L123"></a>pminus1 := new(big.Int).Sub(priv.P, bigOne);
    <a id="L124"></a>qminus1 := new(big.Int).Sub(priv.Q, bigOne);
    <a id="L125"></a>totient := new(big.Int).Mul(pminus1, qminus1);
    <a id="L126"></a>e := big.NewInt(int64(priv.E));
    <a id="L127"></a>gcd := new(big.Int);
    <a id="L128"></a>x := new(big.Int);
    <a id="L129"></a>y := new(big.Int);
    <a id="L130"></a>big.GcdInt(gcd, x, y, totient, e);
    <a id="L131"></a>if gcd.Cmp(bigOne) != 0 {
        <a id="L132"></a>return os.ErrorString(&#34;invalid public exponent E&#34;)
    <a id="L133"></a>}
    <a id="L134"></a><span class="comment">// Check that de ≡ 1 (mod totient(p, q))</span>
    <a id="L135"></a>de := new(big.Int).Mul(priv.D, e);
    <a id="L136"></a>de.Mod(de, totient);
    <a id="L137"></a>if de.Cmp(bigOne) != 0 {
        <a id="L138"></a>return os.ErrorString(&#34;invalid private exponent D&#34;)
    <a id="L139"></a>}
    <a id="L140"></a>return nil;
<a id="L141"></a>}

<a id="L143"></a><span class="comment">// GenerateKeyPair generates an RSA keypair of the given bit size.</span>
<a id="L144"></a>func GenerateKey(rand io.Reader, bits int) (priv *PrivateKey, err os.Error) {
    <a id="L145"></a>priv = new(PrivateKey);
    <a id="L146"></a><span class="comment">// Smaller public exponents lead to faster public key</span>
    <a id="L147"></a><span class="comment">// operations. Since the exponent must be coprime to</span>
    <a id="L148"></a><span class="comment">// (p-1)(q-1), the smallest possible value is 3. Some have</span>
    <a id="L149"></a><span class="comment">// suggested that a larger exponent (often 2**16+1) be used</span>
    <a id="L150"></a><span class="comment">// since previous implementation bugs[1] were avoided when this</span>
    <a id="L151"></a><span class="comment">// was the case. However, there are no current reasons not to use</span>
    <a id="L152"></a><span class="comment">// small exponents.</span>
    <a id="L153"></a><span class="comment">// [1] http://marc.info/?l=cryptography&amp;m=115694833312008&amp;w=2</span>
    <a id="L154"></a>priv.E = 3;

    <a id="L156"></a>pminus1 := new(big.Int);
    <a id="L157"></a>qminus1 := new(big.Int);
    <a id="L158"></a>totient := new(big.Int);

    <a id="L160"></a>for {
        <a id="L161"></a>p, err := randomSafePrime(rand, bits/2);
        <a id="L162"></a>if err != nil {
            <a id="L163"></a>return
        <a id="L164"></a>}

        <a id="L166"></a>q, err := randomSafePrime(rand, bits/2);
        <a id="L167"></a>if err != nil {
            <a id="L168"></a>return
        <a id="L169"></a>}

        <a id="L171"></a>if p.Cmp(q) == 0 {
            <a id="L172"></a>continue
        <a id="L173"></a>}

        <a id="L175"></a>n := new(big.Int).Mul(p, q);
        <a id="L176"></a>pminus1.Sub(p, bigOne);
        <a id="L177"></a>qminus1.Sub(q, bigOne);
        <a id="L178"></a>totient.Mul(pminus1, qminus1);

        <a id="L180"></a>g := new(big.Int);
        <a id="L181"></a>priv.D = new(big.Int);
        <a id="L182"></a>y := new(big.Int);
        <a id="L183"></a>e := big.NewInt(int64(priv.E));
        <a id="L184"></a>big.GcdInt(g, priv.D, y, e, totient);

        <a id="L186"></a>if g.Cmp(bigOne) == 0 {
            <a id="L187"></a>priv.D.Add(priv.D, totient);
            <a id="L188"></a>priv.P = p;
            <a id="L189"></a>priv.Q = q;
            <a id="L190"></a>priv.N = n;

            <a id="L192"></a>break;
        <a id="L193"></a>}
    <a id="L194"></a>}

    <a id="L196"></a>return;
<a id="L197"></a>}

<a id="L199"></a><span class="comment">// incCounter increments a four byte, big-endian counter.</span>
<a id="L200"></a>func incCounter(c *[4]byte) {
    <a id="L201"></a>if c[3]++; c[3] != 0 {
        <a id="L202"></a>return
    <a id="L203"></a>}
    <a id="L204"></a>if c[2]++; c[2] != 0 {
        <a id="L205"></a>return
    <a id="L206"></a>}
    <a id="L207"></a>if c[1]++; c[1] != 0 {
        <a id="L208"></a>return
    <a id="L209"></a>}
    <a id="L210"></a>c[0]++;
<a id="L211"></a>}

<a id="L213"></a><span class="comment">// mgf1XOR XORs the bytes in out with a mask generated using the MGF1 function</span>
<a id="L214"></a><span class="comment">// specified in PKCS#1 v2.1.</span>
<a id="L215"></a>func mgf1XOR(out []byte, hash hash.Hash, seed []byte) {
    <a id="L216"></a>var counter [4]byte;

    <a id="L218"></a>done := 0;
    <a id="L219"></a>for done &lt; len(out) {
        <a id="L220"></a>hash.Write(seed);
        <a id="L221"></a>hash.Write(counter[0:4]);
        <a id="L222"></a>digest := hash.Sum();
        <a id="L223"></a>hash.Reset();

        <a id="L225"></a>for i := 0; i &lt; len(digest) &amp;&amp; done &lt; len(out); i++ {
            <a id="L226"></a>out[done] ^= digest[i];
            <a id="L227"></a>done++;
        <a id="L228"></a>}
        <a id="L229"></a>incCounter(&amp;counter);
    <a id="L230"></a>}
<a id="L231"></a>}

<a id="L233"></a><span class="comment">// MessageTooLongError is returned when attempting to encrypt a message which</span>
<a id="L234"></a><span class="comment">// is too large for the size of the public key.</span>
<a id="L235"></a>type MessageTooLongError struct{}

<a id="L237"></a>func (MessageTooLongError) String() string {
    <a id="L238"></a>return &#34;message too long for RSA public key size&#34;
<a id="L239"></a>}

<a id="L241"></a>func encrypt(c *big.Int, pub *PublicKey, m *big.Int) *big.Int {
    <a id="L242"></a>e := big.NewInt(int64(pub.E));
    <a id="L243"></a>c.Exp(m, e, pub.N);
    <a id="L244"></a>return c;
<a id="L245"></a>}

<a id="L247"></a><span class="comment">// EncryptOAEP encrypts the given message with RSA-OAEP.</span>
<a id="L248"></a><span class="comment">// The message must be no longer than the length of the public modulus less</span>
<a id="L249"></a><span class="comment">// twice the hash length plus 2.</span>
<a id="L250"></a>func EncryptOAEP(hash hash.Hash, rand io.Reader, pub *PublicKey, msg []byte, label []byte) (out []byte, err os.Error) {
    <a id="L251"></a>hash.Reset();
    <a id="L252"></a>k := (pub.N.Len() + 7) / 8;
    <a id="L253"></a>if len(msg) &gt; k-2*hash.Size()-2 {
        <a id="L254"></a>err = MessageTooLongError{};
        <a id="L255"></a>return;
    <a id="L256"></a>}

    <a id="L258"></a>hash.Write(label);
    <a id="L259"></a>lHash := hash.Sum();
    <a id="L260"></a>hash.Reset();

    <a id="L262"></a>em := make([]byte, k);
    <a id="L263"></a>seed := em[1 : 1+hash.Size()];
    <a id="L264"></a>db := em[1+hash.Size() : len(em)];

    <a id="L266"></a>bytes.Copy(db[0:hash.Size()], lHash);
    <a id="L267"></a>db[len(db)-len(msg)-1] = 1;
    <a id="L268"></a>bytes.Copy(db[len(db)-len(msg):len(db)], msg);

    <a id="L270"></a>_, err = io.ReadFull(rand, seed);
    <a id="L271"></a>if err != nil {
        <a id="L272"></a>return
    <a id="L273"></a>}

    <a id="L275"></a>mgf1XOR(db, hash, seed);
    <a id="L276"></a>mgf1XOR(seed, hash, db);

    <a id="L278"></a>m := new(big.Int);
    <a id="L279"></a>m.SetBytes(em);
    <a id="L280"></a>c := encrypt(new(big.Int), pub, m);
    <a id="L281"></a>out = c.Bytes();
    <a id="L282"></a>return;
<a id="L283"></a>}

<a id="L285"></a><span class="comment">// A DecryptionError represents a failure to decrypt a message.</span>
<a id="L286"></a><span class="comment">// It is deliberately vague to avoid adaptive attacks.</span>
<a id="L287"></a>type DecryptionError struct{}

<a id="L289"></a>func (DecryptionError) String() string { return &#34;RSA decryption error&#34; }

<a id="L291"></a><span class="comment">// modInverse returns ia, the inverse of a in the multiplicative group of prime</span>
<a id="L292"></a><span class="comment">// order n. It requires that a be a member of the group (i.e. less than n).</span>
<a id="L293"></a>func modInverse(a, n *big.Int) (ia *big.Int) {
    <a id="L294"></a>g := new(big.Int);
    <a id="L295"></a>x := new(big.Int);
    <a id="L296"></a>y := new(big.Int);
    <a id="L297"></a>big.GcdInt(g, x, y, a, n);
    <a id="L298"></a>if x.Cmp(bigOne) &lt; 0 {
        <a id="L299"></a><span class="comment">// 0 is not the multiplicative inverse of any element so, if x</span>
        <a id="L300"></a><span class="comment">// &lt; 1, then x is negative.</span>
        <a id="L301"></a>x.Add(x, n)
    <a id="L302"></a>}

    <a id="L304"></a>return x;
<a id="L305"></a>}

<a id="L307"></a><span class="comment">// decrypt performs an RSA decryption, resulting in a plaintext integer. If a</span>
<a id="L308"></a><span class="comment">// random source is given, RSA blinding is used.</span>
<a id="L309"></a>func decrypt(rand io.Reader, priv *PrivateKey, c *big.Int) (m *big.Int, err os.Error) {
    <a id="L310"></a><span class="comment">// TODO(agl): can we get away with reusing blinds?</span>
    <a id="L311"></a>if c.Cmp(priv.N) &gt; 0 {
        <a id="L312"></a>err = DecryptionError{};
        <a id="L313"></a>return;
    <a id="L314"></a>}

    <a id="L316"></a>var ir *big.Int;
    <a id="L317"></a>if rand != nil {
        <a id="L318"></a><span class="comment">// Blinding enabled. Blinding involves multiplying c by r^e.</span>
        <a id="L319"></a><span class="comment">// Then the decryption operation performs (m^e * r^e)^d mod n</span>
        <a id="L320"></a><span class="comment">// which equals mr mod n. The factor of r can then be removed</span>
        <a id="L321"></a><span class="comment">// by multipling by the multiplicative inverse of r.</span>

        <a id="L323"></a>r, err1 := randomNumber(rand, priv.N);
        <a id="L324"></a>if err1 != nil {
            <a id="L325"></a>err = err1;
            <a id="L326"></a>return;
        <a id="L327"></a>}
        <a id="L328"></a>if r.Cmp(bigZero) == 0 {
            <a id="L329"></a>r = bigOne
        <a id="L330"></a>}
        <a id="L331"></a>ir = modInverse(r, priv.N);
        <a id="L332"></a>bigE := big.NewInt(int64(priv.E));
        <a id="L333"></a>rpowe := new(big.Int).Exp(r, bigE, priv.N);
        <a id="L334"></a>c.Mul(c, rpowe);
        <a id="L335"></a>c.Mod(c, priv.N);
    <a id="L336"></a>}

    <a id="L338"></a>m = new(big.Int).Exp(c, priv.D, priv.N);

    <a id="L340"></a>if ir != nil {
        <a id="L341"></a><span class="comment">// Unblind.</span>
        <a id="L342"></a>m.Mul(m, ir);
        <a id="L343"></a>m.Mod(m, priv.N);
    <a id="L344"></a>}

    <a id="L346"></a>return;
<a id="L347"></a>}

<a id="L349"></a><span class="comment">// DecryptOAEP decrypts ciphertext using RSA-OAEP.</span>
<a id="L350"></a><span class="comment">// If rand != nil, DecryptOAEP uses RSA blinding to avoid timing side-channel attacks.</span>
<a id="L351"></a>func DecryptOAEP(hash hash.Hash, rand io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) (msg []byte, err os.Error) {
    <a id="L352"></a>k := (priv.N.Len() + 7) / 8;
    <a id="L353"></a>if len(ciphertext) &gt; k ||
        <a id="L354"></a>k &lt; hash.Size()*2+2 {
        <a id="L355"></a>err = DecryptionError{};
        <a id="L356"></a>return;
    <a id="L357"></a>}

    <a id="L359"></a>c := new(big.Int).SetBytes(ciphertext);

    <a id="L361"></a>m, err := decrypt(rand, priv, c);
    <a id="L362"></a>if err != nil {
        <a id="L363"></a>return
    <a id="L364"></a>}

    <a id="L366"></a>hash.Write(label);
    <a id="L367"></a>lHash := hash.Sum();
    <a id="L368"></a>hash.Reset();

    <a id="L370"></a><span class="comment">// Converting the plaintext number to bytes will strip any</span>
    <a id="L371"></a><span class="comment">// leading zeros so we may have to left pad. We do this unconditionally</span>
    <a id="L372"></a><span class="comment">// to avoid leaking timing information. (Although we still probably</span>
    <a id="L373"></a><span class="comment">// leak the number of leading zeros. It&#39;s not clear that we can do</span>
    <a id="L374"></a><span class="comment">// anything about this.)</span>
    <a id="L375"></a>em := leftPad(m.Bytes(), k);

    <a id="L377"></a>firstByteIsZero := subtle.ConstantTimeByteEq(em[0], 0);

    <a id="L379"></a>seed := em[1 : hash.Size()+1];
    <a id="L380"></a>db := em[hash.Size()+1 : len(em)];

    <a id="L382"></a>mgf1XOR(seed, hash, db);
    <a id="L383"></a>mgf1XOR(db, hash, seed);

    <a id="L385"></a>lHash2 := db[0:hash.Size()];

    <a id="L387"></a><span class="comment">// We have to validate the plaintext in contanst time in order to avoid</span>
    <a id="L388"></a><span class="comment">// attacks like: J. Manger. A Chosen Ciphertext Attack on RSA Optimal</span>
    <a id="L389"></a><span class="comment">// Asymmetric Encryption Padding (OAEP) as Standardized in PKCS #1</span>
    <a id="L390"></a><span class="comment">// v2.0. In J. Kilian, editor, Advances in Cryptology.</span>
    <a id="L391"></a>lHash2Good := subtle.ConstantTimeCompare(lHash, lHash2);

    <a id="L393"></a><span class="comment">// The remainder of the plaintext must be zero or more 0x00, followed</span>
    <a id="L394"></a><span class="comment">// by 0x01, followed by the message.</span>
    <a id="L395"></a><span class="comment">//   lookingForIndex: 1 iff we are still looking for the 0x01</span>
    <a id="L396"></a><span class="comment">//   index: the offset of the first 0x01 byte</span>
    <a id="L397"></a><span class="comment">//   invalid: 1 iff we saw a non-zero byte before the 0x01.</span>
    <a id="L398"></a>var lookingForIndex, index, invalid int;
    <a id="L399"></a>lookingForIndex = 1;
    <a id="L400"></a>rest := db[hash.Size():len(db)];

    <a id="L402"></a>for i := 0; i &lt; len(rest); i++ {
        <a id="L403"></a>equals0 := subtle.ConstantTimeByteEq(rest[i], 0);
        <a id="L404"></a>equals1 := subtle.ConstantTimeByteEq(rest[i], 1);
        <a id="L405"></a>index = subtle.ConstantTimeSelect(lookingForIndex&amp;equals1, i, index);
        <a id="L406"></a>lookingForIndex = subtle.ConstantTimeSelect(equals1, 0, lookingForIndex);
        <a id="L407"></a>invalid = subtle.ConstantTimeSelect(lookingForIndex&amp;^equals0, 1, invalid);
    <a id="L408"></a>}

    <a id="L410"></a>if firstByteIsZero&amp;lHash2Good&amp;^invalid&amp;^lookingForIndex != 1 {
        <a id="L411"></a>err = DecryptionError{};
        <a id="L412"></a>return;
    <a id="L413"></a>}

    <a id="L415"></a>msg = rest[index+1 : len(rest)];
    <a id="L416"></a>return;
<a id="L417"></a>}

<a id="L419"></a><span class="comment">// leftPad returns a new slice of length size. The contents of input are right</span>
<a id="L420"></a><span class="comment">// aligned in the new slice.</span>
<a id="L421"></a>func leftPad(input []byte, size int) (out []byte) {
    <a id="L422"></a>n := len(input);
    <a id="L423"></a>if n &gt; size {
        <a id="L424"></a>n = size
    <a id="L425"></a>}
    <a id="L426"></a>out = make([]byte, size);
    <a id="L427"></a>bytes.Copy(out[len(out)-n:len(out)], input);
    <a id="L428"></a>return;
<a id="L429"></a>}
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
