<script>
	import footerLogo from '$lib/images/MetroRod.svg';
	import CCMLetters from '$lib/images/CCMLetters.svg';
	import githubLogo from '$lib/images/github-logo.svg';
	import fire from '$lib/images/fire.svg';
	import linked from '$lib/images/linked.svg';
	import TwitterLogo from '$lib/images/Twitter-Logo.svg';
	import IgLogo from '$lib/images/Ig-logo.svg';
	import kryten from '$lib/images/kryten2.svg';
</script>

<footer class="footer">
	<div class="footer--above">
		<ul>
			<a href="/#body">
				<li>Home</li>
			</a>
			<a href="/about/#body">
				<li>Instrument Registration</li>
			</a>
			<a href="/software/#body">
				<li>Instrument Missing Report</li>
			</a>
			<a href="/embedded/#body">
				<li>About Page</li>
			</a>
			
		</ul>
		<ul>
			<a   href="/embedded/#body">
				<li>Projet Roadmap</li>
			</a>
			<a  data-tooltip="Developer or not, here's how you can contribute" href="/embedded/#body">
				<li >Governance</li>
			</a>
			<a  data-tooltip="Developer or not, here's how you can contribute" href="/embedded/#body">
				<li>Code of Conduct</li>
			</a>
			<a  data-tooltip="Developer or not, here's how you can contribute" href="/embedded/#body">
				<li >How to Contribute</li>
			</a>	
		</ul>

		<ul>
			<a   href="/embedded/#body">
				<li>Contributors(Thank You)</li>
			</a>
			<a  data-tooltip="Developer or not, here's how you can contribute" href="/embedded/#body">
				<li> Software License</li>
			</a>
		</ul>
		
		<img class="building-logo" alt="The project logo" src={footerLogo} />
	</div>
	<div class="footer--iconss">
		<a
			href="https://github.com/Unearthlyglow/sveltego"
			target="_blank"
			data-tooltip="Source Code!"
			class="footer--github"
		>
			<img class="footer--icons" alt="Github Logo" src={githubLogo} />
		</a>
		<a
			href="https://www.instagram.com/ccm_employee/"
			target="_blank"
			data-tooltip="Instagram"
			class="footer--ig"><img class="footer--icons" alt="Instagram Logo" src={IgLogo} /></a
		>

		<!-- <a
			href="https://www.linkedin.com/in/award-malisi/"
			target="_blank"
			data-tooltip="LinkedIn"
			class="footer--linkedin"><img class="footer--icons" alt="LinkedIn Logo" src={linked} /></a
		> -->

		<a
			href="https://twitter.com/CCMAnd2Cents"
			target="_blank"
			data-tooltip="Twitter"
			class="footer--twitter"><img class="footer--icons" alt="Twitter Logo" src={TwitterLogo} /></a
		>
	</div>

	<div class="below--footer">
		<hr class="footer--hr" />

		<div class="footer--ccmquote">
		<div class="footer--images">
			<div class="footer--images--ccmlogo-quote">
			<img class="CCM-letters" src={CCMLetters} alt="CCM Letters" />
			<p class="footer--quote">"It's (Almost) Better Than Life"</p>
		</div>

		<div class="footer--kryten-fire">
		<img class="footer--kryten" src={kryten} alt="Kryten" />

		<img class="footer--fire" src={fire} alt="CCM Letters" />
		</div>
		
		</div>

	

		<p class="footer--Bottomtext">
			Ich bin immer noch dabei, das herauszufinden, also verzeihen Sie mir bitte etwaige Fehler, auf
			die Sie stoßen. Bitte Für Feedback wenden Sie sich bitte an award@capsules.io. Es können alle
			unlustigen Beschwerden eingereicht werden an DoNotReply@capsules.io.
		</p>
		<p class="footer--version">Version: <span>0.0.1</span></p>
		<p class="footer--version">Last Updated: <span>2 / 28 / 24</span></p>
	
</div>
</footer>

<style lang="scss" global>
@import '$lib/styles/global.scss';

	.footer {
		background-color: rgb(244, 245, 225);
		padding-top: 2rem;
	}

	.footer--images {
		display: flex;
		justify-content: space-around;
		align-content: center;
	}

	.footer--images--ccmlogo-quote {
		justify-content: center;
	}

	.footer--above {
		display: flex;
		justify-content: space-around;
		line-height: 1.8rem;
	}

	.footer--above > ul > a > li {
		font-size: 0.7rem;
		color: rgb(13, 81, 90);
		list-style-type: katakana;
	}

	.footer--above > ul > li {
		font-size: 0.7rem;
		color: rgb(13, 81, 90);
		list-style-type: katakana;
	}

	.footer--above > ul > a > li:hover {
		color: #8c2525;
		text-decoration: underline;
	}

	.footer--footloose {
		font-size: 0.7rem;
		font-weight: 900;
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.footer--friends {
		margin: 0.3rem;
		font-size: 0.8rem;
		font-weight: 900;
		color: rgb(13, 81, 90);
	}

	.below--footer {
		margin: 1rem;
	}

	a {
		text-decoration: underline;
		color: #363262;
	}

	a:hover {
		color: #d60675;
	}

	.footer--Bottomtext {
		color: #712828;
		font-weight: 700;
		font-size: 0.8rem;
		margin: 1.5rem;
		margin-top: 3rem;
		// text-align: center;
	}

	.footer--ccmquote {
		display: flex;
		flex-direction: column;
		align-items: space-around;
		border-left: solid 1px rgb(193, 193, 213);
		border-right: solid 1px rgb(193, 193, 213);
		// justify-content: space-between;

	}

	.footer--kryten {
		height: 8rem;
	}

	.footer--quote {
		font-size: 0.8rem;
		font-weight: 800;
		color: rgb(13, 81, 90);
		display: inline-block;
	}


	.footer--kryten-fire {
		display: block;
		justify-content: space-around;

	}

	.building-logo {
		height: 10rem;
	}

	.footer--icons {
		height: 1.3rem;
		margin: 1.3rem;
	}

	.footer--github {
		background: #f7be2f;
	}

	.footer--github:hover {
		background: #fddf92;
	}

	.footer--linkedin {
		background: #363262;
	}

	.footer--linkedin:hover {
		background: #ba91ef;
	}

	.footer--twitter {
		background: #860b0b;
	}

	.footer--twitter:hover {
		background: #dd5050;
	}

	.footer--ig {
		background: #c4c1c1;
	}

	.footer--ig:hover {
		background: #ffffff;
	}

	.footer--version {
		font-size: 0.9rem;
		font-weight: 800;
		color: #363262;
		text-align: right;
		margin: 1rem;
	}

	.footer--version > span {
		color: rgb(13, 81, 90);
	}

	.CCM-letters {
		height: 6rem;
		width: 6rem;
		margin-top: 1rem;
		margin-bottom: 0.5rem;
		
	}

	hr {
		border: 0;
		clear: both;
		display: block;
		width: 96%;
		background-color: #d3a632;
		height: 1px;
	}

	.footer--iconss {
		text-align: center;
		font-size: 3rem;
		padding: 1rem;
	}


	@media only screen and (max-width: 920px) {

.footer--images {
	display: block;
	text-align: center;
}

.footer--quote {
display: block;
}

.footer--kryten-fire {
display: flex;
justify-content: space-around;
}

}

@media only screen and (max-width: 800px) {
		.painting1 {
			max-height: 10rem;
			max-width: 10rem;
		}

		.footer--above {
			margin-left: 0.5rem;
		}

		.footer--icons {
			height: 1.7rem;
		}

		.footer--iconss {
			text-align: center;
			font-size: 2rem;
		}

		.building-logo {
			height: 6rem;
			margin-right: 1.5rem;
			margin-left: 1rem;
		}
	
		.footer--above > ul > a > li {
			font-size: 0.76rem;
		}
	}

@media only screen and (max-width: 548px) {
		.footer--above > ul > li {
			font-size: 0.65rem;
			color: rgb(13, 81, 90);
			list-style-type: katakana;
			margin-left: 1.5rem;
		}

		.footer--above > ul > a > li {
			font-size: 0.65rem;
			color: rgb(13, 81, 90);
	
			margin-left: 1.5rem;
		}

	
	}

@media only screen and (max-width: 450px) {
		.footer--images {
			text-align: center;
	}

	.footer--fire {
	margin-top: 1.5rem;
	}

	.footer--kryten-fire {
		display: block;
	}
}
</style>
