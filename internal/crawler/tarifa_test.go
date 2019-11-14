package crawler

import "testing"

var pagina string = `
<!DOCTYPE html>
<html lang="pt-BR" prefix="og: http://ogp.me/ns#">
<head>
<meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no, minimal-ui">
<meta name="google-site-verification" content="v7Xly3sJZhBWskzoWD33NLOOdDbqvyXXIRbg0lSWg6U" />
<link rel="profile" href="http://gmpg.org/xfn/11">
<title>Página inicial | SmartMEI - O Anjo da Guarda do MEI</title>

<!-- This site is optimized with the Yoast SEO plugin v7.9 - https://yoast.com/wordpress/plugins/seo/ -->
<link rel="canonical" href="https://www.smartmei.com.br/" />
<meta property="og:locale" content="pt_BR" />
<meta property="og:type" content="website" />
<meta property="og:title" content="Página inicial | SmartMEI - O Anjo da Guarda do MEI" />
<meta property="og:url" content="https://www.smartmei.com.br/" />
<meta property="og:site_name" content="SmartMEI - O Anjo da Guarda do MEI" />
<meta name="twitter:card" content="summary" />
<meta name="twitter:title" content="Página inicial | SmartMEI - O Anjo da Guarda do MEI" />
<script type='application/ld+json'>{"@context":"https:\/\/schema.org","@type":"WebSite","@id":"#website","url":"https:\/\/www.smartmei.com.br\/","name":"SmartMEI - O Anjo da Guarda do MEI","potentialAction":{"@type":"SearchAction","target":"https:\/\/www.smartmei.com.br\/?s={search_term_string}","query-input":"required name=search_term_string"}}</script>
<script type='application/ld+json'>{"@context":"https:\/\/schema.org","@type":"Organization","url":"https:\/\/www.smartmei.com.br\/","sameAs":["https:\/\/www.facebook.com\/smartmeiapp","https:\/\/www.linkedin.com\/company-beta\/10107529","https:\/\/www.youtube.com\/channel\/UCIlNUVHjyLhsZrLGX-gCumQ"],"@id":"https:\/\/www.smartmei.com.br\/#organization","name":"SmartMEI","logo":"https:\/\/www.smartmei.com.br\/wp-content\/uploads\/2017\/09\/cropped-icon6-tab-SM-3.png"}</script>
<!-- / Yoast SEO plugin. -->

<link rel='dns-prefetch' href='//ajax.googleapis.com' />
<link rel='dns-prefetch' href='//maxcdn.bootstrapcdn.com' />
<link rel='dns-prefetch' href='//fonts.googleapis.com' />
<link rel='dns-prefetch' href='//s.w.org' />
<link rel="alternate" type="application/rss+xml" title="Feed para SmartMEI - O Anjo da Guarda do MEI &raquo;" href="https://www.smartmei.com.br/feed/" />
<!-- This site uses the Google Analytics by MonsterInsights plugin v7.10.0 - Using Analytics tracking - https://www.monsterinsights.com/ -->
<script type="text/javascript" data-cfasync="false">
	var mi_version         = '7.10.0';
	var mi_track_user      = true;
	var mi_no_track_reason = '';
	
	var disableStr = 'ga-disable-UA-36901076-3';

	/* Function to detect opted out users */
	function __gaTrackerIsOptedOut() {
		return document.cookie.indexOf(disableStr + '=true') > -1;
	}

	/* Disable tracking if the opt-out cookie exists. */
	if ( __gaTrackerIsOptedOut() ) {
		window[disableStr] = true;
	}

	/* Opt-out function */
	function __gaTrackerOptout() {
	  document.cookie = disableStr + '=true; expires=Thu, 31 Dec 2099 23:59:59 UTC; path=/';
	  window[disableStr] = true;
	}
	
	if ( mi_track_user ) {
		(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
			(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
			m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
		})(window,document,'script','//www.google-analytics.com/analytics.js','__gaTracker');

		__gaTracker('create', 'UA-36901076-3', 'auto');
		__gaTracker('set', 'forceSSL', true);
		__gaTracker('require', 'displayfeatures');
		__gaTracker('send','pageview');
	} else {
		console.log( "" );
		(function() {
			/* https://developers.google.com/analytics/devguides/collection/analyticsjs/ */
			var noopfn = function() {
				return null;
			};
			var noopnullfn = function() {
				return null;
			};
			var Tracker = function() {
				return null;
			};
			var p = Tracker.prototype;
			p.get = noopfn;
			p.set = noopfn;
			p.send = noopfn;
			var __gaTracker = function() {
				var len = arguments.length;
				if ( len === 0 ) {
					return;
				}
				var f = arguments[len-1];
				if ( typeof f !== 'object' || f === null || typeof f.hitCallback !== 'function' ) {
					console.log( 'Not running function __gaTracker(' + arguments[0] + " ....) because you are not being tracked. " + mi_no_track_reason );
					return;
				}
				try {
					f.hitCallback();
				} catch (ex) {

				}
			};
			__gaTracker.create = function() {
				return new Tracker();
			};
			__gaTracker.getByName = noopnullfn;
			__gaTracker.getAll = function() {
				return [];
			};
			__gaTracker.remove = noopfn;
			window['__gaTracker'] = __gaTracker;
					})();
		}
</script>
<!-- / Google Analytics by MonsterInsights -->
		<script type="text/javascript">
			window._wpemojiSettings = {"baseUrl":"https:\/\/s.w.org\/images\/core\/emoji\/11\/72x72\/","ext":".png","svgUrl":"https:\/\/s.w.org\/images\/core\/emoji\/11\/svg\/","svgExt":".svg","source":{"concatemoji":"https:\/\/www.smartmei.com.br\/wp-includes\/js\/wp-emoji-release.min.js?ver=5.0.7"}};
			!function(a,b,c){function d(a,b){var c=String.fromCharCode;l.clearRect(0,0,k.width,k.height),l.fillText(c.apply(this,a),0,0);var d=k.toDataURL();l.clearRect(0,0,k.width,k.height),l.fillText(c.apply(this,b),0,0);var e=k.toDataURL();return d===e}function e(a){var b;if(!l||!l.fillText)return!1;switch(l.textBaseline="top",l.font="600 32px Arial",a){case"flag":return!(b=d([55356,56826,55356,56819],[55356,56826,8203,55356,56819]))&&(b=d([55356,57332,56128,56423,56128,56418,56128,56421,56128,56430,56128,56423,56128,56447],[55356,57332,8203,56128,56423,8203,56128,56418,8203,56128,56421,8203,56128,56430,8203,56128,56423,8203,56128,56447]),!b);case"emoji":return b=d([55358,56760,9792,65039],[55358,56760,8203,9792,65039]),!b}return!1}function f(a){var c=b.createElement("script");c.src=a,c.defer=c.type="text/javascript",b.getElementsByTagName("head")[0].appendChild(c)}var g,h,i,j,k=b.createElement("canvas"),l=k.getContext&&k.getContext("2d");for(j=Array("flag","emoji"),c.supports={everything:!0,everythingExceptFlag:!0},i=0;i<j.length;i++)c.supports[j[i]]=e(j[i]),c.supports.everything=c.supports.everything&&c.supports[j[i]],"flag"!==j[i]&&(c.supports.everythingExceptFlag=c.supports.everythingExceptFlag&&c.supports[j[i]]);c.supports.everythingExceptFlag=c.supports.everythingExceptFlag&&!c.supports.flag,c.DOMReady=!1,c.readyCallback=function(){c.DOMReady=!0},c.supports.everything||(h=function(){c.readyCallback()},b.addEventListener?(b.addEventListener("DOMContentLoaded",h,!1),a.addEventListener("load",h,!1)):(a.attachEvent("onload",h),b.attachEvent("onreadystatechange",function(){"complete"===b.readyState&&c.readyCallback()})),g=c.source||{},g.concatemoji?f(g.concatemoji):g.wpemoji&&g.twemoji&&(f(g.twemoji),f(g.wpemoji)))}(window,document,window._wpemojiSettings);
		</script>
		<style type="text/css">
img.wp-smiley,
img.emoji {
	display: inline !important;
	border: none !important;
	box-shadow: none !important;
	height: 1em !important;
	width: 1em !important;
	margin: 0 .07em !important;
	vertical-align: -0.1em !important;
	background: none !important;
	padding: 0 !important;
}
</style>
<link rel='stylesheet' id='wp-block-library-css'  href='https://www.smartmei.com.br/wp-includes/css/dist/block-library/style.min.css?ver=5.0.7' type='text/css' media='all' />
<link rel='stylesheet' id='fvp-frontend-css'  href='https://www.smartmei.com.br/wp-content/plugins/featured-video-plus/styles/frontend.css?ver=2.3.3' type='text/css' media='all' />
<link rel='stylesheet' id='taxonomy-image-plugin-public-css'  href='https://www.smartmei.com.br/wp-content/plugins/taxonomy-images/css/style.css?ver=0.9.6' type='text/css' media='screen' />
<link rel='stylesheet' id='woocommerce-layout-css'  href='https://www.smartmei.com.br/wp-content/plugins/woocommerce/assets/css/woocommerce-layout.css?ver=3.5.2' type='text/css' media='all' />
<style id='woocommerce-layout-inline-css' type='text/css'>

	.infinite-scroll .woocommerce-pagination {
		display: none;
	}
</style>
<link rel='stylesheet' id='woocommerce-smallscreen-css'  href='https://www.smartmei.com.br/wp-content/plugins/woocommerce/assets/css/woocommerce-smallscreen.css?ver=3.5.2' type='text/css' media='only screen and (max-width: 768px)' />
<link rel='stylesheet' id='woocommerce-general-css'  href='https://www.smartmei.com.br/wp-content/plugins/woocommerce/assets/css/woocommerce.css?ver=3.5.2' type='text/css' media='all' />
<style id='woocommerce-inline-inline-css' type='text/css'>
.woocommerce form .form-row .required { visibility: visible; }
</style>
<link rel='stylesheet' id='wordpress-popular-posts-css-css'  href='https://www.smartmei.com.br/wp-content/plugins/wordpress-popular-posts/public/css/wpp.css?ver=4.0.13' type='text/css' media='all' />
<link rel='stylesheet' id='smartmei_2017-style-google-fonts-css'  href='https://fonts.googleapis.com/css?family=Open+Sans:400,600|Raleway:400,500,600' type='text/css' media='all' />
<link rel='stylesheet' id='smartmei_2017-style-bootstrap-css'  href='https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='smartmei_2017-style-fontawesome-css'  href='https://www.smartmei.com.br/wp-content/themes/smartmei_2017/css/font-awesome.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='smartmei_2017-style-css'  href='https://www.smartmei.com.br/wp-content/themes/smartmei_2017/style.css?ver=201802231520' type='text/css' media='all' />
<link rel='stylesheet' id='smartmei_2017-new-styles-css'  href='https://www.smartmei.com.br/wp-content/themes/smartmei_2017/css/new-styles.css' type='text/css' media='all' />
<link rel='stylesheet' id='tablepress-default-css'  href='https://www.smartmei.com.br/wp-content/plugins/tablepress/css/default.min.css?ver=1.9' type='text/css' media='all' />
<link rel='stylesheet' id='elementor-icons-css'  href='https://www.smartmei.com.br/wp-content/plugins/elementor/assets/lib/eicons/css/elementor-icons.min.css?ver=1.9.3' type='text/css' media='all' />
<link rel='stylesheet' id='font-awesome-css'  href='https://www.smartmei.com.br/wp-content/plugins/elementor/assets/lib/font-awesome/css/font-awesome.min.css?ver=4.7.0' type='text/css' media='all' />
<link rel='stylesheet' id='elementor-animations-css'  href='https://www.smartmei.com.br/wp-content/plugins/elementor/assets/css/animations.min.css?ver=1.9.3' type='text/css' media='all' />
<link rel='stylesheet' id='elementor-frontend-css'  href='https://www.smartmei.com.br/wp-content/plugins/elementor/assets/css/frontend.min.css?ver=1.9.3' type='text/css' media='all' />
<link rel='stylesheet' id='elementor-global-css'  href='https://www.smartmei.com.br/wp-content/uploads/elementor/css/global.css?ver=1564418913' type='text/css' media='all' />
<script type='text/javascript'>
/* <![CDATA[ */
var monsterinsights_frontend = {"js_events_tracking":"true","download_extensions":"doc,pdf,ppt,zip,xls,docx,pptx,xlsx","inbound_paths":"[]","home_url":"https:\/\/www.smartmei.com.br","hash_tracking":"false"};
/* ]]> */
</script>
<script type='text/javascript' src='https://www.smartmei.com.br/wp-content/plugins/google-analytics-for-wordpress/assets/js/frontend.min.js?ver=7.10.0'></script>
<script type='text/javascript' src='https://www.smartmei.com.br/wp-includes/js/jquery/jquery.js?ver=1.12.4'></script>
<script type='text/javascript' src='https://www.smartmei.com.br/wp-includes/js/jquery/jquery-migrate.min.js?ver=1.4.1'></script>
<script type='text/javascript' src='https://www.smartmei.com.br/wp-content/plugins/featured-video-plus/js/jquery.fitvids.min.js?ver=master-2015-08'></script>
<script type='text/javascript'>
/* <![CDATA[ */
var fvpdata = {"ajaxurl":"https:\/\/www.smartmei.com.br\/wp-admin\/admin-ajax.php","nonce":"46d00f1eb0","fitvids":"1","dynamic":"","overlay":"","opacity":"0.75","color":"b","width":"640"};
/* ]]> */
</script>
<script type='text/javascript' src='https://www.smartmei.com.br/wp-content/plugins/featured-video-plus/js/frontend.min.js?ver=2.3.3'></script>
<link rel='https://api.w.org/' href='https://www.smartmei.com.br/wp-json/' />
<link rel="EditURI" type="application/rsd+xml" title="RSD" href="https://www.smartmei.com.br/xmlrpc.php?rsd" />
<link rel="wlwmanifest" type="application/wlwmanifest+xml" href="https://www.smartmei.com.br/wp-includes/wlwmanifest.xml" /> 
<meta name="generator" content="WordPress 5.0.7" />
<meta name="generator" content="WooCommerce 3.5.2" />
<link rel='shortlink' href='https://www.smartmei.com.br/' />
<link rel="alternate" type="application/json+oembed" href="https://www.smartmei.com.br/wp-json/oembed/1.0/embed?url=https%3A%2F%2Fwww.smartmei.com.br%2F" />
         <style type="text/css">
             #banner { 
			 	background-image: url(https://www.smartmei.com.br/wp-content/uploads/2017/03/bg-banner.jpg); 
			}
         </style>
    	<noscript><style>.woocommerce-product-gallery{ opacity: 1 !important; }</style></noscript>
		<style type="text/css">
			.site-title,
		.site-description {
			position: absolute;
			clip: rect(1px, 1px, 1px, 1px);
		}
		</style>
	<link rel="stylesheet" type="text/css" href="https://fonts.googleapis.com/css?family=Roboto:100,100italic,200,200italic,300,300italic,400,400italic,500,500italic,600,600italic,700,700italic,800,800italic,900,900italic%7CRoboto+Slab:100,100italic,200,200italic,300,300italic,400,400italic,500,500italic,600,600italic,700,700italic,800,800italic,900,900italic"><link rel="icon" href="https://www.smartmei.com.br/wp-content/uploads/2017/09/cropped-icon6-tab-SM-3-32x32.png" sizes="32x32" />
<link rel="icon" href="https://www.smartmei.com.br/wp-content/uploads/2017/09/cropped-icon6-tab-SM-3-192x192.png" sizes="192x192" />
<link rel="apple-touch-icon-precomposed" href="https://www.smartmei.com.br/wp-content/uploads/2017/09/cropped-icon6-tab-SM-3-180x180.png" />
<meta name="msapplication-TileImage" content="https://www.smartmei.com.br/wp-content/uploads/2017/09/cropped-icon6-tab-SM-3-270x270.png" />
		<style type="text/css" id="wp-custom-css">
			#billing_country_field {
	display: none;
}		</style>
		
<!-- Hotjar Tracking Code for www.smartmei.com.br -->
<script>
  (function(h,o,t,j,a,r){
      h.hj=h.hj||function(){(h.hj.q=h.hj.q||[]).push(arguments)};
      h._hjSettings={hjid:1118281,hjsv:6};
      a=o.getElementsByTagName('head')[0];
      r=o.createElement('script');r.async=1;
      r.src=t+h._hjSettings.hjid+j+h._hjSettings.hjsv;
      a.appendChild(r);
  })(window,document,'https://static.hotjar.com/c/hotjar-','.js?sv=');
</script>
</head>

<body class="home page-template-default page page-id-10 wp-custom-logo woocommerce-no-js group-blog elementor-default elementor-page elementor-page-10" data-spy="scroll" data-target="#navbar-menu">

<nav class="navbar navbar-smartmei navbar-fixed-top" id="navbar-menu" role="navigation">
	<div class="content-container">
		<!-- Brand and toggle get grouped for better mobile display -->
		<div class="navbar-header col-sm-3 col-xs-12">
			<button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#navbar-mobile" aria-expanded="false">
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
			</button>

							<a class="navbar-brand hidden-xs" href="https://www.smartmei.com.br/">
					<img src="https://www.smartmei.com.br/wp-content/uploads/2017/04/LogoSite.png" alt="SmartMEI" class="img-responsive"/>
				</a>
				<a class="navbar-brand visible-xs" href="https://www.smartmei.com.br/">
					<img src="https://www.smartmei.com.br/wp-content/uploads/2017/04/LogoSite.png" alt="SmartMEI" height="35"/>
				</a>
					
		</div>

		<!-- Collect the nav links, forms, and other content for toggling -->
		<div class="collapse navbar-collapse col-sm-9 col-xs-12 pull-right" id="navbar-mobile">
			<ul class="nav navbar-nav navbar-right">
				<li class="visible-xs clearfix">
					<button type="button" class="btn btn-link btn-close-menu pull-right" data-toggle="collapse" data-target="#navbar-mobile">
						&times;
					</button>
				</li>
				<li id="menu-item-21" class="visible-xs menu-item menu-item-type-custom menu-item-object-custom current-menu-item current_page_item menu-item-21"><a title="Home" href="/">HOME</a></li>
<li id="menu-item-87" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-87"><a title="Como Funciona" href="#burocracia">COMO FUNCIONA</a></li>
<li id="menu-item-88" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-88"><a title="Planos e Tarifas" href="#planos-e-tarifas">PLANOS E TARIFAS</a></li>
<li id="menu-item-2647" class="menu-item menu-item-type-post_type menu-item-object-page menu-item-2647"><a title="Central de Ajuda" href="https://www.smartmei.com.br/central-de-ajuda/">Central de ajuda</a></li>
				<li>
					<a class="btn btn-default" href="https://play.google.com/store/apps/details?id=br.com.smartmei&hl=pt_BR" target="_blank">
						BAIXAR APP
					</a>
				</li>
			</ul>
										<div class="menu-support visible-xs">
					<h5>
						<span>SUPORTE</span>
						<hr/>
					</h5>
					<p class="text-center">
						Seg - Sex:  9h às 18h					</p>
				</div>
					</div>
	</div>
</nav>

	<header id="banner">
		<div class="content-container">
			<h1>
				Ser MEI ficou fácil de verdade!			</h1>
			<h4>
				Tudo o que você precisa para nunca mais se preocupar com governo, contadores ou bancos.			</h4>
							<div class="banner-link">
					<a class="btn btn-lg btn-download" href="https://play.google.com/store/apps/details?id=br.com.smartmei&hl=pt_BR" target="_blank">
						BAIXAR APP
					</a>
				</div>
					</div>
		<div class="content-container banner-container">
			<div class="banner-circle">
				<picture>
					<source media="(max-width: 767px)" srcset="https://www.smartmei.com.br/wp-content/themes/smartmei_2017/img/bg-circle-mobile.png">
					<source media="(min-width: 768px)" srcset="https://www.smartmei.com.br/wp-content/themes/smartmei_2017/img/bg-circle.png">
					<img class="img-responsive" srcset="https://www.smartmei.com.br/wp-content/themes/smartmei_2017/img/bg-circle.png">
				</picture>
				<div class="banner-animation">
					<picture>
						<source media="(max-width: 767px)" srcset="https://www.smartmei.com.br/wp-content/themes/smartmei_2017/img/banner_15fps_320px.gif 320w, https://www.smartmei.com.br/wp-content/themes/smartmei_2017/img/banner_15fps_600px.gif 600w">
						<source media="(min-width: 768px)" srcset="https://www.smartmei.com.br/wp-content/themes/smartmei_2017/img/banner_15fps_75.gif 900w, https://www.smartmei.com.br/wp-content/themes/smartmei_2017/img/banner_15fps_66.gif 792w">
						<img class="img-responsive" srcset="https://www.smartmei.com.br/wp-content/themes/smartmei_2017/img/banner_15fps_75.gif">
					</picture>
				</div>
				<div class="banner-phone">
					<img class="img-responsive" src="https://www.smartmei.com.br/wp-content/themes/smartmei_2017/img/banner-phone.png">
				</div>
			</div>
		</div>
		<div class="banner-bottom">
			<picture>
				<source media="(max-width: 767px)" srcset="https://www.smartmei.com.br/wp-content/themes/smartmei_2017/img/banner-mask-mobile.png">
				<source media="(min-width: 768px)" srcset="https://www.smartmei.com.br/wp-content/themes/smartmei_2017/img/banner-mask.png">
				<img srcset="https://www.smartmei.com.br/wp-content/themes/smartmei_2017/img/banner-mask.png">
			</picture>
		</div>
	</header>
		
		

<section id="burocracia" style="background-image: url();">
	<div class="content-container"> 
		<div class="row">
			<div class="col-sm-10 col-sm-offset-1">
				<h2>Fique livre da burocracia</h2>
<h4>Cuidamos de tudo para sua empresa estar sempre em dia com o governo.</h4>
	<div id="smartmeibc_801" class="carousel slide" data-ride="carousel">
		<!-- Wrapper for slides -->
		<div class="carousel-inner" role="listbox">
							<div class="item active">
					<div class="card-container">
						<div class="card">
							<div class="row clearfix">
								<div class="col-sm-4 carousel-image">
									<img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/bureaucracy-taxes.jpg" class="img-responsive">
								</div>
								<div class="col-sm-8 carousel-content">
									<h3>Controle seus impostos</h3>
									<p>Saiba se os seus impostos (DAS) estão em dia, o valor acumulado das parcelas atrasadas e gere todas as guias de pagamento pelo app.</p>
								</div>
							</div>
						</div>
					</div>                        
				</div>
							<div class="item ">
					<div class="card-container">
						<div class="card">
							<div class="row clearfix">
								<div class="col-sm-4 carousel-image">
									<img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/bureaucracy-tax-return.jpg" class="img-responsive">
								</div>
								<div class="col-sm-8 carousel-content">
									<h3>Faça sua Declaração Anual</h3>
									<p>Faça suas Declarações Anuais (DASN) pelo seu telefone em poucos minutos e sem complicação, inclusive as atrasadas.</p>
								</div>
							</div>
						</div>
					</div>                        
				</div>
							<div class="item ">
					<div class="card-container">
						<div class="card">
							<div class="row clearfix">
								<div class="col-sm-4 carousel-image">
									<img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/bureaucracy-support.jpg" class="img-responsive">
								</div>
								<div class="col-sm-8 carousel-content">
									<h3>Receba suporte personalizado</h3>
									<p>Fale com o nosso time de especialistas pelo chat para esclarecer qualquer pergunta sobre a sua empresa.</p>
								</div>
							</div>
						</div>
					</div>                        
				</div>
							<div class="item ">
					<div class="card-container">
						<div class="card">
							<div class="row clearfix">
								<div class="col-sm-4 carousel-image">
									<img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/bureaucracy-company-data.jpg" class="img-responsive">
								</div>
								<div class="col-sm-8 carousel-content">
									<h3>Atualize seus dados</h3>
									<p>Cheque e atualize os dados cadastrais da sua empresa sem precisar entrar em nenhum site do governo.</p>
								</div>
							</div>
						</div>
					</div>                        
				</div>
					</div>                
		<!-- Controls -->
		<a class="left carousel-control" href="#smartmeibc_801" onclick="this.blur();" role="button" data-slide="prev">
			<span class="fa fa-angle-left" aria-hidden="true"></span>
			<span class="sr-only">Previous</span>
		</a>
		<a class="right carousel-control" href="#smartmeibc_801" onclick="this.blur();" role="button" data-slide="next">
			<span class="fa fa-angle-right" aria-hidden="true"></span>
			<span class="sr-only">Next</span>
		</a>
		<!-- Indicators -->
		<ol class="carousel-indicators">
							<li data-target="#smartmeibc_801" data-slide-to="0" class="active"></li>
							<li data-target="#smartmeibc_801" data-slide-to="1" class=""></li>
							<li data-target="#smartmeibc_801" data-slide-to="2" class=""></li>
							<li data-target="#smartmeibc_801" data-slide-to="3" class=""></li>
					</ol>
	</div>
    
			</div>
		</div>
	</div>
</section>

<section id="cobranca" style="background-image: url();">
	<div class="content-container"> 
		<div class="row">
			<div class="col-sm-10 col-sm-offset-1">
				<h2>Receba pagamentos em nome da sua empresa</h2>
<h4>Cobre de maneira profissional e seja bem visto pelos seus clientes.</h4>
	<div id="smartmeibc_103" class="carousel slide" data-ride="carousel">
		<!-- Wrapper for slides -->
		<div class="carousel-inner" role="listbox">
							<div class="item active">
					<div class="card-container">
						<div class="card">
							<div class="row clearfix">
								<div class="col-sm-4 carousel-image">
									<img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/billing-banking.jpg" class="img-responsive">
								</div>
								<div class="col-sm-8 carousel-content">
									<h3>Parecido com banco, mas sem papelada</h3>
									<p>Uma conta virtual para a sua empresa, aprovada na hora e sem precisar preencher nada.</p>
								</div>
							</div>
						</div>
					</div>                        
				</div>
							<div class="item ">
					<div class="card-container">
						<div class="card">
							<div class="row clearfix">
								<div class="col-sm-4 carousel-image">
									<img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/billing-billet.jpg" class="img-responsive">
								</div>
								<div class="col-sm-8 carousel-content">
									<h3>Boletos sem complicação</h3>
									<p>Envie boletos com o nome e a marca da sua empresa e receba em até dois dias úteis após o pagamento.</p>
								</div>
							</div>
						</div>
					</div>                        
				</div>
							<div class="item ">
					<div class="card-container">
						<div class="card">
							<div class="row clearfix">
								<div class="col-sm-4 carousel-image">
									<img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/billing-credit-card.jpg" class="img-responsive">
								</div>
								<div class="col-sm-8 carousel-content">
									<h3>Pagamentos com cartão de crédito e débito</h3>
									<p>Contrate uma maquininha para que seus clientes possam te pagar usando cartão de crédito e débito.&gt;
<h4><a href="https://www.smartmei.com.br/sobre-a-maquininha/">NOVIDADE!</a></h4></p>
								</div>
							</div>
						</div>
					</div>                        
				</div>
							<div class="item ">
					<div class="card-container">
						<div class="card">
							<div class="row clearfix">
								<div class="col-sm-4 carousel-image">
									<img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/billing-deposit.jpg" class="img-responsive">
								</div>
								<div class="col-sm-8 carousel-content">
									<h3>Receba depósitos de empresas parceiras</h3>
									<p>Use sua conta para receber pagamentos de empresas parceiras da SmartMEI.
<h4>LANÇAMENTO EM BREVE</h4></p>
								</div>
							</div>
						</div>
					</div>                        
				</div>
							<div class="item ">
					<div class="card-container">
						<div class="card">
							<div class="row clearfix">
								<div class="col-sm-4 carousel-image">
									<img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/billing-fees.jpg" class="img-responsive">
								</div>
								<div class="col-sm-8 carousel-content">
									<h3>Preços justos e sem tarifas fixas</h3>
									<p>Condições pensadas especialmente para a realidade do MEI.
<h4><a href="#planos-e-tarifas">VEJA AS NOSSAS TARIFAS ABAIXO</a></h4></p>
								</div>
							</div>
						</div>
					</div>                        
				</div>
					</div>                
		<!-- Controls -->
		<a class="left carousel-control" href="#smartmeibc_103" onclick="this.blur();" role="button" data-slide="prev">
			<span class="fa fa-angle-left" aria-hidden="true"></span>
			<span class="sr-only">Previous</span>
		</a>
		<a class="right carousel-control" href="#smartmeibc_103" onclick="this.blur();" role="button" data-slide="next">
			<span class="fa fa-angle-right" aria-hidden="true"></span>
			<span class="sr-only">Next</span>
		</a>
		<!-- Indicators -->
		<ol class="carousel-indicators">
							<li data-target="#smartmeibc_103" data-slide-to="0" class="active"></li>
							<li data-target="#smartmeibc_103" data-slide-to="1" class=""></li>
							<li data-target="#smartmeibc_103" data-slide-to="2" class=""></li>
							<li data-target="#smartmeibc_103" data-slide-to="3" class=""></li>
							<li data-target="#smartmeibc_103" data-slide-to="4" class=""></li>
					</ol>
	</div>
    
			</div>
		</div>
	</div>
</section>

<section id="controle-financeiro" style="background-image: url();">
	<div class="content-container"> 
		<div class="row">
			<div class="col-sm-10 col-sm-offset-1">
				<h2>Controle Financeiro mais simples</h2>
<h4>Um jeito mais fácil de cuidar do seu dinheiro e de saber se o seu negócio está dando retorno.</h4>
	<div id="smartmeibc_900" class="carousel slide" data-ride="carousel">
		<!-- Wrapper for slides -->
		<div class="carousel-inner" role="listbox">
							<div class="item active">
					<div class="card-container">
						<div class="card">
							<div class="row clearfix">
								<div class="col-sm-4 carousel-image">
									<img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/finance-reports.jpg" class="img-responsive">
								</div>
								<div class="col-sm-8 carousel-content">
									<h3>Controle sem digitar nada</h3>
									<p>Receba dos seus clientes e pague seu fornecedores pela SmartMEI e saiba seu lucro automaticamente.</p>
								</div>
							</div>
						</div>
					</div>                        
				</div>
							<div class="item ">
					<div class="card-container">
						<div class="card">
							<div class="row clearfix">
								<div class="col-sm-4 carousel-image">
									<img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/finance-card.jpg" class="img-responsive">
								</div>
								<div class="col-sm-8 carousel-content">
									<h3>Faça saques e compras no débito</h3>
									<p>Você recebe um cartão de débito VISA SmartMEI para usar seu saldo com ainda mais conveniência.</p>
								</div>
							</div>
						</div>
					</div>                        
				</div>
							<div class="item ">
					<div class="card-container">
						<div class="card">
							<div class="row clearfix">
								<div class="col-sm-4 carousel-image">
									<img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/finance-transfer.jpg" class="img-responsive">
								</div>
								<div class="col-sm-8 carousel-content">
									<h3>Transfira dinheiro e pague contas pelo app</h3>
									<p>Faça transferências para qualquer conta em banco e pague boletos com código de barras.</p>
								</div>
							</div>
						</div>
					</div>                        
				</div>
					</div>                
		<!-- Controls -->
		<a class="left carousel-control" href="#smartmeibc_900" onclick="this.blur();" role="button" data-slide="prev">
			<span class="fa fa-angle-left" aria-hidden="true"></span>
			<span class="sr-only">Previous</span>
		</a>
		<a class="right carousel-control" href="#smartmeibc_900" onclick="this.blur();" role="button" data-slide="next">
			<span class="fa fa-angle-right" aria-hidden="true"></span>
			<span class="sr-only">Next</span>
		</a>
		<!-- Indicators -->
		<ol class="carousel-indicators">
							<li data-target="#smartmeibc_900" data-slide-to="0" class="active"></li>
							<li data-target="#smartmeibc_900" data-slide-to="1" class=""></li>
							<li data-target="#smartmeibc_900" data-slide-to="2" class=""></li>
					</ol>
	</div>
    
			</div>
		</div>
	</div>
</section>

<section id="ainda-nao-tem-mei" style="background-image: url(https://www.smartmei.com.br/wp-content/uploads/2017/03/bg-mei.jpg);">
	<div class="content-container"> 
		<div class="row">
			<div class="col-sm-10 col-sm-offset-1">
				<div class="row clearfix">
<div class="col-xs-12">
<h2>Ainda não tem MEI? Não tem problema!</h2>
<h4>Faça a abertura pelo nosso app e receba seu <strong>CNPJ</strong> em um dia útil. <strong>É de graça!</strong></h4>
</div>
</div>
<div class="row clearfix"><a class="col-md-3 col-sm-4 col-xs-6" href="https://play.google.com/store/apps/details?id=br.com.smartmei&amp;hl=pt_BR" target="_blank"><br />
<img class="img-responsive" src="/wp-content/uploads/2017/03/google-play-button-pt-300x90.png" alt="Disponível na Google Play" /></a></div>
			</div>
		</div>
	</div>
</section>

<section id="planos-e-tarifas" style="background-image: url();">
	<div class="content-container"> 
		<div class="row">
			<div class="col-sm-10 col-sm-offset-1">
				<h2>Conheça nossos planos e tarifas!</h2>
<h4>Compare a SmartMEI com o custo de um contador e de uma conta Pessoa Jurídica nos principais bancos</h4>
<ul class="nav nav-tabs" role="tablist"><li role="presentation" class="active"><a href="#planos" aria-controls="planos" role="tab" data-toggle="tab">Planos</a></li><li role="presentation" class=""><a href="#tarifas-2" aria-controls="tarifas-2" role="tab" data-toggle="tab">Tarifas</a></li><li role="presentation" class=""><a href="#compare" aria-controls="compare" role="tab" data-toggle="tab">Compare</a></li></ul><div class="tab-content"><div role="tabpanel" class="tab-pane table-plans active" id="planos">
	                                    <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            <h3>CONTABILIDADE</h3>
                        </div>
                        <div class="text-center col-sm-3 col-xs-6 planos-0-1">
                            <h4>PLANO<br/>Básico</h4>                        </div>
                        <div class="text-center col-sm-3 col-xs-6 planos-0-2">
                            <h4>PLANO<br/>Profissional</h4>
                        </div>
                                        
                </div>
            
                            <div class="row visible-xs">
                    <div class="col-xs-12 cell-title">
                        <h3>CONTABILIDADE</h3>
                    </div>
                </div>
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Emissão de guias de pagamento de imposto                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-1-1">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Emissão de guias de pagamento de imposto                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-1-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Verificar situação de pagamento de impostos                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-2-1">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Verificar situação de pagamento de impostos                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-2-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Fazer declaração Anual                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-3-1">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Fazer declaração Anual                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-3-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Suporte de time de especialistas                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-4-1">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Suporte de time de especialistas                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-4-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Entrada de vendas e despesas                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-5-1">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Entrada de vendas e despesas                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-5-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Contabilização sem esforço de despesas                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-6-1">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Contabilização sem esforço de despesas                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-6-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Relatório de lucro da MEI                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-7-1">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Relatório de lucro da MEI                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-7-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                                                            <div class="row">
                    <div class="col-xs-12 cell-title">
                        <h3>SERVIÇOS FINANCEIROS
</h3>
                    </div>
                </div>
            
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Recebimento de parceiros                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-9-1">
                                                    </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Recebimento de parceiros                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-9-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Receber pagamento via boleto                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-10-1">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Receber pagamento via boleto                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-10-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Transferência para conta bancária do dono da MEI                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-11-1">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Transferência para conta bancária do dono da MEI                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-11-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Transferência para qualquer conta bancária                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-12-1">
                                                    </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Transferência para qualquer conta bancária                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-12-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Saque                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-13-1">
                                                    </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Saque                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-13-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Pagamento de contas com código de barra                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-14-1">
                                                    </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Pagamento de contas com código de barra                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-14-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Compras com cartão de débito Visa SmartMEI                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-15-1">
                                                    </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Compras com cartão de débito Visa SmartMEI                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-15-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                                                            <div class="row row-eq-height">
                                            <div class="col-sm-6 hidden-xs">
                            Realizar vendas por maquininha                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-16-1">
                                                    </div>
                        <div class="col-sm-6 col-xs-8 text-center visible-xs">
                            Realizar vendas por maquininha                        </div>
                        <div class="text-center col-sm-3 col-xs-2 planos-16-2">
                            <span class="fa fa-check item-available-pro" aria-hidden="true"></span>                        </div>
                                        
                </div>
            
                        </div><div role="tabpanel" class="tab-pane table-plans" id="tarifas-2">
	                    <div class="row row-eq-height">
                <div class="col-sm-4 hidden-xs">
                                    </div>
                <div class="text-center col-sm-4 col-xs-6 tarifas-2-0-1">
                    <h4>PLANO<br>Básico</h4>                </div>
                <div class="text-center col-sm-4 col-xs-6 tarifas-2-0-2">
                    <h4>PLANO<br>Profissional</h4>                </div>
            </div>
                                <div class="row row-eq-height">
                <div class="col-sm-4 cell-small-title">
                    Cobrança                </div>
                <div class="text-center col-sm-4 col-xs-6 tarifas-2-1-1">
                    R$ 5,00 por boleto pago.                </div>
                <div class="text-center col-sm-4 col-xs-6 tarifas-2-1-2">
                    R$ 5,00 por boleto pago.                </div>
            </div>
                                <div class="row row-eq-height">
                <div class="col-sm-4 cell-small-title">
                    Transferência                </div>
                <div class="text-center col-sm-4 col-xs-6 tarifas-2-2-1">
                    R$ 7,00*<br/>
*Limitado a contas PF do dono da MEI                </div>
                <div class="text-center col-sm-4 col-xs-6 tarifas-2-2-2">
                    R$ 7,00                </div>
            </div>
                                <div class="row row-eq-height">
                <div class="col-sm-4 cell-small-title">
                    Saque                </div>
                <div class="text-center col-sm-4 col-xs-6 tarifas-2-3-1">
                    --                </div>
                <div class="text-center col-sm-4 col-xs-6 tarifas-2-3-2">
                    R$ 7,00
                </div>
            </div>
                                <div class="row row-eq-height">
                <div class="col-sm-4 cell-small-title">
                    Mensalidade                </div>
                <div class="text-center col-sm-4 col-xs-6 tarifas-2-4-1">
                    --                </div>
                <div class="text-center col-sm-4 col-xs-6 tarifas-2-4-2">
                    R$ 15,00<br/>
*pagando R$45,00/trimestre                </div>
            </div>
                                <div class="row row-eq-height">
                <div class="col-sm-4 cell-small-title">
                    Vendas por maquininha                </div>
                <div class="text-center col-sm-4 col-xs-6 tarifas-2-5-1">
                    --                </div>
                <div class="text-center col-sm-4 col-xs-6 tarifas-2-5-2">
                    <a href="https://www.smartmei.com.br/sobre-a-maquininha/#taxas">Veja mais</a>                </div>
            </div>
            </div><div role="tabpanel" class="tab-pane table-plans" id="compare">
	                    
                            <div class="row">
                    <div class="text-center col-sm-2 col-sm-offset-8 col-xs-6">
                        <img class="img-responsive" src="/wp-content/uploads/2017/04/LogoFooter.png">                    </div>
                    <div class="text-center col-sm-2 col-xs-6">
                                                <div class="btn-group btn-img-select">
                            <button class="btn btn-default btn-lg dropdown-toggle" type="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                                                    <img id="bank_img_bb" class=" bank-image" src="https://s3-sa-east-1.amazonaws.com/smartmei-static/banks/bb.png" class="img-responsive">
                                                                    <img id="bank_img_bradesco" class="hidden bank-image" src="https://s3-sa-east-1.amazonaws.com/smartmei-static/banks/bradesco.png" class="img-responsive">
                                                                    <img id="bank_img_caixa" class="hidden bank-image" src="https://s3-sa-east-1.amazonaws.com/smartmei-static/banks/caixa.png" class="img-responsive">
                                                                    <img id="bank_img_itau" class="hidden bank-image" src="https://s3-sa-east-1.amazonaws.com/smartmei-static/banks/itau.png" class="img-responsive">
                                                                    <img id="bank_img_santander" class="hidden bank-image" src="https://s3-sa-east-1.amazonaws.com/smartmei-static/banks/santander.png" class="img-responsive">
                                                                <span class="fa fa-sort-desc" aria-hidden="true"></span>
                            </button>
                            <ul class="dropdown-menu">
                                                                    <li class="bank_option_bb">
                                        <a href="javascript:;" data-toggle="select" data-image="bank_img_bb" data-display="bank-option-bb">
                                            <img src="https://s3-sa-east-1.amazonaws.com/smartmei-static/banks/bb.png" class="img-responsive">
                                        </a>
                                    </li>
                                                                    <li class="bank_option_bradesco">
                                        <a href="javascript:;" data-toggle="select" data-image="bank_img_bradesco" data-display="bank-option-bradesco">
                                            <img src="https://s3-sa-east-1.amazonaws.com/smartmei-static/banks/bradesco.png" class="img-responsive">
                                        </a>
                                    </li>
                                                                    <li class="bank_option_caixa">
                                        <a href="javascript:;" data-toggle="select" data-image="bank_img_caixa" data-display="bank-option-caixa">
                                            <img src="https://s3-sa-east-1.amazonaws.com/smartmei-static/banks/caixa.png" class="img-responsive">
                                        </a>
                                    </li>
                                                                    <li class="bank_option_itau">
                                        <a href="javascript:;" data-toggle="select" data-image="bank_img_itau" data-display="bank-option-itau">
                                            <img src="https://s3-sa-east-1.amazonaws.com/smartmei-static/banks/itau.png" class="img-responsive">
                                        </a>
                                    </li>
                                                                    <li class="bank_option_santander">
                                        <a href="javascript:;" data-toggle="select" data-image="bank_img_santander" data-display="bank-option-santander">
                                            <img src="https://s3-sa-east-1.amazonaws.com/smartmei-static/banks/santander.png" class="img-responsive">
                                        </a>
                                    </li>
                                                            </ul>
                        </div>
                    </div>
                </div>
                                            
                                                <div class="row row-eq-height">
                        <div class="col-sm-8 col-xs-12 cell-small-title">
                            Saldo e Extrato Online                        </div>
                        <div class="text-center cell-gray col-sm-2 col-xs-6">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option  bank-option-bb">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-bradesco">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-caixa">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-itau">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-santander">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                            </div>
                                                            
                                                <div class="row row-eq-height">
                        <div class="col-sm-8 col-xs-12 cell-small-title">
                            Saques em Caixas Eletrônicos                        </div>
                        <div class="text-center cell-gray col-sm-2 col-xs-6">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option  bank-option-bb">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-bradesco">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-caixa">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-itau">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-santander">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                            </div>
                                                            
                                                <div class="row row-eq-height">
                        <div class="col-sm-8 col-xs-12 cell-small-title">
                            Pagamento de Boletos pelo Celular                        </div>
                        <div class="text-center cell-gray col-sm-2 col-xs-6">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option  bank-option-bb">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-bradesco">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-caixa">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-itau">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-santander">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                            </div>
                                                            
                                                <div class="row row-eq-height">
                        <div class="col-sm-8 col-xs-12 cell-small-title">
                            Cartão para Movimentar a Conta                        </div>
                        <div class="text-center cell-gray col-sm-2 col-xs-6">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option  bank-option-bb">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-bradesco">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-caixa">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-itau">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-santander">
                                <span class="fa fa-check item-available" aria-hidden="true"></span>                            </div>
                                            </div>
                                                            
                                                <div class="row row-eq-height">
                        <div class="col-sm-8 col-xs-12 cell-small-title">
                            Aprovação Imediata e 100% Garantida                        </div>
                        <div class="text-center cell-gray col-sm-2 col-xs-6">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option  bank-option-bb">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-bradesco">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-caixa">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-itau">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-santander">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                            </div>
                                                            
                                                <div class="row row-eq-height">
                        <div class="col-sm-8 col-xs-12 cell-small-title">
                            Boletos com a Marca de sua Empresa                        </div>
                        <div class="text-center cell-gray col-sm-2 col-xs-6">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option  bank-option-bb">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-bradesco">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-caixa">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-itau">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-santander">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                            </div>
                                                            
                                                <div class="row row-eq-height">
                        <div class="col-sm-8 col-xs-12 cell-small-title">
                            Apoio contábil                        </div>
                        <div class="text-center cell-gray col-sm-2 col-xs-6">
                            <span class="fa fa-check item-available" aria-hidden="true"></span>                        </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option  bank-option-bb">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-bradesco">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-caixa">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-itau">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-santander">
                                <span class="fa fa-times item-unavailable" aria-hidden="true"></span>                            </div>
                                            </div>
                                                            
                                                <div class="row">
                        <div class="col-xs-12">
                            &nbsp;
                        </div>
                    </div>
                                                            
                                                <div class="row row-eq-height">
                        <div class="col-sm-8 col-xs-12 cell-small-title">
                            Conta PJ – Mensalidade                        </div>
                        <div class="text-center cell-gray col-sm-2 col-xs-6">
                            R$ 15,00                        </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option  bank-option-bb">
                                R$ 18,00                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-bradesco">
                                R$ 43,70                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-caixa">
                                R$ 35,00                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-itau">
                                R$ 47,00                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-santander">
                                R$ 28,50                            </div>
                                            </div>
                                                            
                                                <div class="row row-eq-height">
                        <div class="col-sm-8 col-xs-12 cell-small-title">
                            Preço por boleto pago                        </div>
                        <div class="text-center cell-gray col-sm-2 col-xs-6">
                            R$ 5,00                        </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option  bank-option-bb">
                                R$ 9,00                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-bradesco">
                                R$ 11,00                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-caixa">
                                R$ 5,90                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-itau">
                                R$ 10,00                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-santander">
                                R$ 12,00                            </div>
                                            </div>
                                                            
                                                <div class="row row-eq-height">
                        <div class="col-sm-8 col-xs-12 cell-small-title">
                            Preço por transferência                        </div>
                        <div class="text-center cell-gray col-sm-2 col-xs-6">
                            R$ 7,00                        </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option  bank-option-bb">
                                R$ 10,00                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-bradesco">
                                R$ 8,95                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-caixa">
                                R$ 8,65                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-itau">
                                R$ 9,55                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-santander">
                                R$ 8,80                            </div>
                                            </div>
                                                            
                                                <div class="row">
                        <div class="col-xs-12">
                            &nbsp;
                        </div>
                    </div>
                                                            
                                                <div class="row row-eq-height">
                        <div class="col-sm-8 col-xs-12 cell-small-title">
                            Preço de Contador                        </div>
                        <div class="text-center cell-gray col-sm-2 col-xs-6">
                            R$ 0,00<br/>
&nbsp;                        </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option  bank-option-bb">
                                R$ 100,00*<br/><small>* preço médio</small>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-bradesco">
                                R$ 100,00*<br/><small>* preço médio</small>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-caixa">
                                R$ 100,00*<br/><small>* preço médio</small>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-itau">
                                R$ 100,00*<br/><small>* preço médio</small>                            </div>
                                                    <div class="text-center col-sm-2 col-xs-6 bank-option hidden bank-option-santander">
                                R$ 100,00*<br/><small>* preço médio</small>                            </div>
                                            </div>
                                                            
                            <div class="row">
                    <div class="col-xs-12">
                        &nbsp;
                    </div>
                </div>
                <div class="row row-footer row-eq-height hidden-xs">
                    <div class="col-sm-8">
                        <span>TOTAL da mensalidade<br/>CONTA PJ + CONTADOR</span>                    </div>
                    <div class="text-center cell-gray col-sm-2">
                        R$ 15,00<br/>*R$ 45,00/trimestre                    </div>
                                            <div class="text-center col-sm-2 last-column bank-option  bank-option-bb">
                            R$ 118,00                        </div>
                                            <div class="text-center col-sm-2 last-column bank-option hidden bank-option-bradesco">
                            R$ 143,70
                        </div>
                                            <div class="text-center col-sm-2 last-column bank-option hidden bank-option-caixa">
                            R$ 135,00                        </div>
                                            <div class="text-center col-sm-2 last-column bank-option hidden bank-option-itau">
                            R$ 147,00
                        </div>
                                            <div class="text-center col-sm-2 last-column bank-option hidden bank-option-santander">
                            R$ 128,50                        </div>
                                    </div>
                <div class="row row-footer visible-xs">
                    <div class="col-sm-8 col-xs-12">
                        <span>TOTAL da mensalidade<br/>CONTA PJ + CONTADOR</span>                    </div>
                    <div class="text-center cell-gray col-sm-2 col-xs-6">
                        R$ 15,00<br/>*R$ 45,00/trimestre                    </div>
                                            <div class="text-center col-sm-2 last-column col-xs-6 bank-option  bank-option-bb">
                            R$ 118,00                        </div>
                                            <div class="text-center col-sm-2 last-column col-xs-6 bank-option hidden bank-option-bradesco">
                            R$ 143,70
                        </div>
                                            <div class="text-center col-sm-2 last-column col-xs-6 bank-option hidden bank-option-caixa">
                            R$ 135,00                        </div>
                                            <div class="text-center col-sm-2 last-column col-xs-6 bank-option hidden bank-option-itau">
                            R$ 147,00
                        </div>
                                            <div class="text-center col-sm-2 last-column col-xs-6 bank-option hidden bank-option-santander">
                            R$ 128,50                        </div>
                                    </div>
                            <div class="row no-border">
        <div class="col-xs-12 text-right">
            <a href="/wp-content/uploads/2017/04/tabela-comparativa.pdf" target="_blank" title="VER COMPARAÇÃO COM TODOS OS BANCOS">
                VER COMPARAÇÃO COM TODOS OS BANCOS
            </a>
        </div>
    </div>
</div></div>
			</div>
		</div>
	</div>
</section>

<section id="depoimentos" style="background-image: url();">
	<div class="content-container"> 
		<div class="row">
			<div class="col-sm-10 col-sm-offset-1">
				<h2><strong>Depoimentos</strong> de pessoas felizes</h2>
<p>&nbsp;</p>
<div class="row clearfix row-users row-eq-height"><div class="col-sm-4 box-user text-center">
    <div class="box-content">
        <div class="box-thumb">
            <img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/photo.jpg" class="img-circle">
        </div>
        <div class="box-title">
            Vivian Scherlloch        </div>
        <div class="box-location">
            Vitória        </div>
        <div class="box-rating">
                            <span class="fa fa-star" aria-hidden="true"></span>
                            <span class="fa fa-star" aria-hidden="true"></span>
                            <span class="fa fa-star" aria-hidden="true"></span>
                            <span class="fa fa-star" aria-hidden="true"></span>
                            <span class="fa fa-star" aria-hidden="true"></span>
                    </div>
        <p class="box-testimonial">
            <p>O aplicativo é maravilhoso, atende todas as minhas necessidades inclusive me manteve a par de informações que eu nem me recordava em atualizar&#8230; Já havia comentado e escrevi ótimo&#8230; Mas após analisar mais a fundo fiquei impressionada com tantos pontos positivos em um APP &#8230; Amei</p>
        </p>
    </div>
</div><div class="col-sm-4 box-user text-center">
    <div class="box-content">
        <div class="box-thumb">
            <img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/photo-user-1.jpg" class="img-circle">
        </div>
        <div class="box-title">
            Flavio Martins        </div>
        <div class="box-location">
            São Paulo        </div>
        <div class="box-rating">
                            <span class="fa fa-star" aria-hidden="true"></span>
                            <span class="fa fa-star" aria-hidden="true"></span>
                            <span class="fa fa-star" aria-hidden="true"></span>
                            <span class="fa fa-star" aria-hidden="true"></span>
                            <span class="fa fa-star" aria-hidden="true"></span>
                    </div>
        <p class="box-testimonial">
            <p>Funciona e a empresa está cada vez mais empenhada em melhorar o APP. Suporte total!!!</p>
        </p>
    </div>
</div><div class="col-sm-4 box-user text-center">
    <div class="box-content">
        <div class="box-thumb">
            <img src="https://www.smartmei.com.br/wp-content/uploads/2017/03/photo-1.jpg" class="img-circle">
        </div>
        <div class="box-title">
            Andreia Erse        </div>
        <div class="box-location">
            Pirassununga        </div>
        <div class="box-rating">
                            <span class="fa fa-star" aria-hidden="true"></span>
                            <span class="fa fa-star" aria-hidden="true"></span>
                            <span class="fa fa-star" aria-hidden="true"></span>
                            <span class="fa fa-star" aria-hidden="true"></span>
                            <span class="fa fa-star" aria-hidden="true"></span>
                    </div>
        <p class="box-testimonial">
            <p>Estou gostando da experiência com o APP. Estou cogitando a conta PJ também&#8230; Uma coisa excelente seria uma maquininha de cartão (débito e credito) integrada SMARTMEI com taxas atraentes ao publico MEI. Fica a sugestão.</p>
        </p>
    </div>
</div></div>
			</div>
		</div>
	</div>
</section>

<section id="cadastre-se" style="background-image: url(https://www.smartmei.com.br/wp-content/uploads/2017/03/bg-signup.jpg);">
	<div class="content-container"> 
		<div class="row">
			<div class="col-sm-10 col-sm-offset-1">
				<h2>Deixe a SmartMEI cuidar da sua empresa.</h2>
<h4>Cadastre-se em menos de um minuto e fique livre de qualquer preocupação.</h4>
<p><a class="btn btn-lg btn-default" href="https://play.google.com/store/apps/details?id=br.com.smartmei&amp;hl=pt_BR" target="_blank">BAIXAR APP</a></p>
			</div>
		</div>
	</div>
</section>
<!-- <section>
	<div class="content-container">
		<div class="accordions">
			<div class="accordion accordion-nested">
				<div class="accordion-head">Sobre o MEI</div>
				<div class="accordion-content">
					<div class="accordion">
						<div class="accordion-head">Sobre o MEI</div>
						<div class="accordion-content">
							Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation. Lorem ipsum dolor sit amet, consectetur adipisicing.
						</div>
					</div>
					<div class="accordion">
						<div class="accordion-head">Sobre o MEI</div>
						<div class="accordion-content">
							Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation. Lorem ipsum dolor sit amet, consectetur adipisicing.
						</div>
					</div>
					<div class="accordion">
						<div class="accordion-head">Sobre o MEI</div>
						<div class="accordion-content">
							Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation. Lorem ipsum dolor sit amet, consectetur adipisicing.
						</div>
					</div>
				</div>
			</div>
			<div class="accordion">
				<div class="accordion-head">Sobre o MEI</div>
				<div class="accordion-content">
					Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation. Lorem ipsum dolor sit amet, consectetur adipisicing.
				</div>
			</div>
			<div class="accordion">
				<div class="accordion-head">Sobre o MEI</div>
				<div class="accordion-content">
					Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation. Lorem ipsum dolor sit amet, consectetur adipisicing.
				</div>
			</div>
		</div>
	</div>
</section>
 -->


<footer>
	<div class="content-container">
		<div class="row">
			<div class="col-md-2 col-xs-12">
				<img class="img-responsive img-logo" src="https://www.smartmei.com.br/wp-content/uploads/2017/04/LogoFooter.png">
			</div>
			
			<div class="col-md-9 col-md-offset-1 col-xs-12">
				<ul class="list-unstyled list-inline list-bottom-menu visible-sm visible-xs">
					<li> 
						© SmartMEI 2016. Todos os Direitos Reservados.					</li>
				</ul>
				<ul class="list-unstyled list-inline list-bottom-menu hidden-sm hidden-xs pull-left">
					<li> 
						© SmartMEI 2016. Todos os Direitos Reservados.					</li>
				</ul>
				<ul class="list-unstyled list-inline list-bottom-menu">
					<li id="menu-item-122" class="menu-item menu-item-type-post_type menu-item-object-page menu-item-122"><a href="https://www.smartmei.com.br/termos-de-uso/">Termos de Uso</a></li>
<li id="menu-item-2764" class="menu-item menu-item-type-post_type menu-item-object-page menu-item-2764"><a href="https://www.smartmei.com.br/politica-de-privacidade/">Política de Privacidade</a></li>
				</ul>
			</div>
		</div>
	</div>
</footer>
<!-- ATIVADO PELO PAINEL -->
<!-- Google Analytics -->
<!-- <script>
(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','https://www.google-analytics.com/analytics.js','ga');

ga('create', 'UA-36901076-3', 'auto');
ga('send', 'pageview');
</script> -->
<!-- End Google Analytics -->
<!-- Google Code for CTR baixar app Conversion Page -->
<script type="text/javascript">
/* <![CDATA[ */
var google_conversion_id = 827810299;
var google_conversion_label = "eLYxCMzJiHkQ-8PdigM";
var google_remarketing_only = false;
/* ]]> */
</script>
<script type="text/javascript" src="//www.googleadservices.com/pagead/conversion.js">
</script>
<noscript>
<div style="display:inline;">
<img height="1" width="1" style="border-style:none;" alt="" src="//www.googleadservices.com/pagead/conversion/827810299/?label=eLYxCMzJiHkQ-8PdigM&amp;guid=ON&amp;script=0"/>
</div>
</noscript>
<!-- End Google Code for CTR baixar app Conversion Page -->

<script type='text/javascript' async src='https://d335luupugsy2.cloudfront.net/js/loader-scripts/9d215d4c-d590-439e-bce4-cdb490ee1c9f-loader.js'></script>	<script type="text/javascript">
		var c = document.body.className;
		c = c.replace(/woocommerce-no-js/, 'woocommerce-js');
		document.body.className = c;
	</script>
	<script type='text/javascript' src='https://www.smartmei.com.br/wp-content/plugins/woocommerce/assets/js/jquery-blockui/jquery.blockUI.min.js?ver=2.70'></script>
<script type='text/javascript'>
/* <![CDATA[ */
var wc_add_to_cart_params = {"ajax_url":"\/wp-admin\/admin-ajax.php","wc_ajax_url":"\/?wc-ajax=%%endpoint%%","i18n_view_cart":"Ver carrinho","cart_url":"https:\/\/www.smartmei.com.br\/cart\/","is_cart":"","cart_redirect_after_add":"no"};
/* ]]> */
</script>
<script type='text/javascript' src='https://www.smartmei.com.br/wp-content/plugins/woocommerce/assets/js/frontend/add-to-cart.min.js?ver=3.5.2'></script>
<script type='text/javascript' src='https://www.smartmei.com.br/wp-content/plugins/woocommerce/assets/js/js-cookie/js.cookie.min.js?ver=2.1.4'></script>
<script type='text/javascript'>
/* <![CDATA[ */
var woocommerce_params = {"ajax_url":"\/wp-admin\/admin-ajax.php","wc_ajax_url":"\/?wc-ajax=%%endpoint%%"};
/* ]]> */
</script>
<script type='text/javascript' src='https://www.smartmei.com.br/wp-content/plugins/woocommerce/assets/js/frontend/woocommerce.min.js?ver=3.5.2'></script>
<script type='text/javascript'>
/* <![CDATA[ */
var wc_cart_fragments_params = {"ajax_url":"\/wp-admin\/admin-ajax.php","wc_ajax_url":"\/?wc-ajax=%%endpoint%%","cart_hash_key":"wc_cart_hash_d24ac26e0df1d2777fd7605a3d88cb7a","fragment_name":"wc_fragments_d24ac26e0df1d2777fd7605a3d88cb7a"};
/* ]]> */
</script>
<script type='text/javascript' src='https://www.smartmei.com.br/wp-content/plugins/woocommerce/assets/js/frontend/cart-fragments.min.js?ver=3.5.2'></script>
<script type='text/javascript' src='https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js'></script>
<script type='text/javascript' src='https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js'></script>
<script type='text/javascript' src='https://www.smartmei.com.br/wp-content/themes/smartmei_2017/js/jquery.bcSwipe.min.js'></script>
<script type='text/javascript' src='https://www.smartmei.com.br/wp-content/themes/smartmei_2017/js/site.js?ver=201802231520'></script>
<script type='text/javascript' src='https://www.smartmei.com.br/wp-includes/js/wp-embed.min.js?ver=5.0.7'></script>

</body>
</html>
`

func TestRecuperarTarifa(t *testing.T) {

	descricaoEsperada := "R$ 7,00"

	descricao := recuperarTarifa(pagina)

	if descricaoEsperada != descricao {
		t.Errorf("Descrição esperada %q, mas recebida %q\n", descricaoEsperada, descricao)
	}

}

func TestConsultarTarifa(t *testing.T) {

	var err error

	valorEsperado := float64(7)
	descricaoEsperada := "R$ 7,00"
	tarifa, err := ConsultarTarifa()

	if err != nil {
		t.Error(err)
		return
	}

	if tarifa == nil {
		t.Errorf("Tarifa nula\n")
		return
	}

	if tarifa != nil && (tarifa.Valor == nil || valorEsperado != *tarifa.Valor) {
		t.Errorf("Valor esperado %v, mas recebido %v\n", valorEsperado, tarifa.Valor)
	}
	if descricaoEsperada != tarifa.Descricao {
		t.Errorf("Descrição esperada %q, mas recebida %q\n", descricaoEsperada, tarifa.Descricao)
	}

}
