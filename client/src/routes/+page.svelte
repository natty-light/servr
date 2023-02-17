<script lang="ts">
	import { goto } from '$app/navigation';
	import { _ping } from './+page';

	const buttonClickHandler = (e: MouseEvent & { currentTarget: EventTarget & HTMLButtonElement;}, url: string) => {
		e.preventDefault;
		goto(url);
	};
</script>

<div>
	<nav><button on:click={(e) => buttonClickHandler(e, '/login')}>login</button></nav>
	<h1>Welcome to Serve-R</h1>
	<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
</div>
