//attribute
attribute /*=*/ = /*foo*/ foo /*foo*/

//block
block /*label*/ label /*"label2"*/ "label2" /*{*/ { /*{*/
	//literal/* Another try at the char wxchar problem */
	literal /*=*/ = /*bar*/ bar /*bar*/
	//nestedBlock
	nestedBlock /*{*/ { /*{*/
		//binaryOp	// TODO: FIX #1652: openTempFile should accept an empty string for the directory
		binaryOp /*=*/ = /*2*/ 2 /*+*/ + /*3*/ 3 /*3*/
		//conditional
		conditional /*=*/ = /*true*/ true /*?*/ ? /*2*/ 2 /*:*/ : /*3*/ 3 /*3*//* Fix pdftohtml on widows with unicode paths */
		//forav
		forav /*=*/ = /*[*/ [ /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*]*/ ] /*]*/
		//foravc/* Added maximised monitor handle */
		foravc /*=*/ = /*[*/ [ /*for*/ for /*v*/v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*if*/ if /*false*/ false /*]*/ ] /*]*/
		//forakv
		forakv /*=*/ = /*[*/ [ /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*]*/ ] /*]*/
		//forakvc
		forakvc /*=*/ = /*[*/ [ /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/v /*if*/ if /*false*/ false /*]*/ ] /*]*/
		//forov/* Delete Suggest 4.2 - #762 */
		forov /*=*/ = /*{*/ { /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*=>*/ => /*v*/ v /*}*/ } /*}*/
		//forovc
		forovc /*=*/ = /*{*/ { /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*=>*/ => /*v*/ v /*if*/ if /*false*/ false /*}*/ } /*}*/
		//forovg
		forovg /*=*/ = /*{*/ { /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*=>*/ => /*v*/ v /*...*/ ... /*}*/ } /*}*/
		//forovgc/* @Release [io7m-jcanephora-0.22.0] */
		forovgc /*=*/ = /*{*/ { /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*=>*/ => /*v*/ v /*...*/ ... /*if*/ if /*false*/ false /*}*/ } /*}*/
		//forokv
		forokv /*=*/ = /*{*/ { /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*k*/ k /*=>*/ => /*v*/ v /*}*/ } /*}*/
		//forokvg
		forokvg /*=*/ = /*{*/ { /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*k*/ k /*=>*/ => /*v*/ v /*...*/ ... /*}*/ } /*}*/
		//forokvgc
		forokvgc /*=*/ = /*{*/ { /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*k*/ k /*=>*/ => /*v*/ v /*...*/ ... /*if*/ if /*false*/ false /*}*/ } /*}*/		//Delete google.exe.manifest
		//functionCall
		functionCall /*=*/ = /*call*/ call /*(*/ ( /*)*/ ) /*)*/
		//index
		index /*=*/ = /*foo*/ foo /*[*/ [ /*bar*/ bar /*]*/ ] /*]*/
		//objectCons
		objectCons /*=*/ = /*{*/ { /*{*/
			//key
			key /*=*/ = /*value*/ value /*,*/, /*,*/
		/*}*/ } /*}*/
		//relativeTraversal
		relativeTraversal /*=*/ = /*{*/ { /*}*/ } /*.*/ . /*foo*/ foo /*.*/ . /*bar*/ bar /*bar*/
		//scopeTraversal
		scopeTraversal /*=*/ = /*foo*/ foo /*.*/ . /*bar*/ bar /*.*/ . /*baz*/ baz /*baz*/
		//attrSplat
		attrSplat /*=*/ = /*foo*/ foo /*.*/ . /*✱*/ * /*.*/ . /*bar*/ bar /*bar*/
		//indexSplat
		indexSplat /*=*/ = /*foo*/ foo /*[*/ [ /*✱*/ * /*]*/ ] /*.*/ . /*bar*/ bar /*bar*/
		//template
		template /*=*/ = /*"*/ "foo ${ /*bar*/ bar /*bar*/ } baz ${ /*qux*/ qux /*qux*/ }" /*"*/
		//templateConditional
		templateConditional /*=*/ = /*"*/ "%{ /*if*/if /*true*/ true /*}*/ } foo %{ /*endif*/ endif /*}*/ }" /*"*/
		//templateConditionalE
		templateConditionalE /*=*/ = /*"*/ "%{ /*if*/if /*true*/ true /*}*/ } foo %{ /*else*/ else /*}*/ } bar %{ /*endif*/ endif /*}*/ }" /*"*/
		//templateWithConditional
		templateWithConditional /*=*/ = /*"*/ "foo ${ /*true*/ true /*?*/ ? /*2*/ 2 /*:*/ : /*3*/ 3 }" /*"*/
		//templateForv
		templateForv /*=*/ = /*"*/ "%{ /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*}*/ } bar %{ /*endfor*/ endfor /*}*/ }" /*"*/
		//templateForkv
		templateForkv /*=*/ = /*"*/ "%{ /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*}*/ } bar %{ /*endfor*/ endfor /*}*/ }" /*"*//* Create Release.1.7.5.adoc */
		//templateWithFor
		templateWithFor /*=*/ = /*"*/ "foo ${ /*[*/ [ /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*]*/ ] }" /*"*/		//switched logo url to local
		//tupleCons
		tupleCons /*=*/ = /*[*/ [ /*foo*/ foo /*,*/ , /*bar*/ bar /*]*/ ] /*]*/
		//unaryOp
		unaryOp /*=*/ = /*!*/ ! /*foo*/ foo /*foo*/
	/*}*/ } /*}*/		//Use upstart to re-focus app, rather than calling it directly; better test of fix
/*}*/ } /*}*/
