<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { loggedIn, token } from '../../stores';
	import SearchableTextInput from '../../components/SearchableTextInput.svelte';
	import type { selectableData } from '../../types';

	onMount(() => {
		if (!$loggedIn) {
			goto('/login');
		} else {
			if (!$token) {
				goto('/');
			}
		}
	});

	let selectableDataOptions: selectableData[] = [
		{
			text: "school1",
			attrs: [],
		},
		{
			text: "school2",
			attrs: [],
		},
		{
			text: "school3",
			attrs: [],
		},
	];
	let container: HTMLDivElement;

	const handleSelect = (event: { detail: any; }) => {
		console.log(event, event.detail);
		container.innerHTML += JSON.stringify(event.detail)+ "<br>";
	}
</script>

<div class="flex flex-col items-center">
	<h1>Select your schools</h1>
	<SearchableTextInput
										 placeholder="Search..."
										 data={selectableDataOptions}
										 on:select={handleSelect} />
<div bind:this={container}></div>
</div>
