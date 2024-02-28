<nav class="container-tabs">
	<li>
		<a href="/#body" class="single--tab"> Home</a>
	</li>
	<li class="class-divide">||</li>
	<li role="list" dir="ltr">
		<p class="tab-intact" aria-haspopup="listbox">Instrument Functions</p>
		<ul role="listbox">
			<li class="manifesto--test"><a href="/jazz-manifesto/#body">Jazz Manifesto</a></li>

			<li>
				<a href="https://unearthlyglow.github.io/RustMDBooks-CCMBlog/" target="_blank"
					>CCM General Blog(MD Book)</a
				>
			</li>
			
		</ul>
		
	</li>
	<li class="class-divide">||</li>

	<li class="class-divide">||</li>
	<li role="list" dir="rtl">
		<p class="tab-intact" aria-haspopup="listbox">Documentation</p>
		<ul role="listbox">
			<li><a href="/softwaremap/#body">Project ReadMe</a></li>
			<li><a href="/embedded/#body">Project Roadmap</a></li>
			<li><a href="/visual/#body">Governance</a></li>
			<li><a href="/art/#body">How to Contribute</a></li>
			<li><a href="/embedded/#body">Code of Conduct</a></li>
			<li><a href="/embedded/#body">Contributors(Thank You)</a></li>
			<li><a href="/embedded/#body">Software License</a></li>
			
		</ul>
	</li>
	<li class="class-divide">||</li>
	<li role="list" dir="rtl">
		<p class="tab-intact" aria-haspopup="listbox">About</p>
		<ul role="listbox">
			<li><a href="/about/#body">About The Project</a></li>
		</ul>
	</li>
</nav>

<style lang="scss">
	@import '$lib/styles/global.scss';

	nav {
		
		color: white;
	}

	.single--tab {
		font-weight: 900;
	}

	.container-tabs {
		justify-content: center;

		& a {
			color: $primary-white;
			font-size: 1.5rem;
		}
	}

	.tab-intact {
		color: $primary-white;
		font-weight: 900;
		margin: 0;
		font-size: 1.5rem;
	}

	.tab-intact:hover {
		color: #942bef;
	}

	.class-divide {
		color: #8c2525;
	}

	ul[role='listbox'] {
		background-color: $primary-white;
		box-shadow: 10px 5px 5px rgba(0, 0, 0, 0.98);
		border: none;
		text-align: center;

		& a {
			font-size: 0.87rem;
			color: $primary-green;
		}
	}

	ul[role='listbox']:hover {
		background-color: $primary-white;
		transition: background-color 0.5s;

		& a {
			color: $primary-green;
		}

		& a:hover {
			text-decoration: underline;
			color: $primary-yellow;
		}
	}

	a:hover {
		color: #d60675;
	}

	@media only screen and (max-width: 548px) {
		.container-tabs {
			display: none;
		}
	}
</style>
