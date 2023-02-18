<!-- Credit to https://svelte.dev/repl/864898d1defc44538db7cb2ee9a5ac97?version=3.25.0 -->

<style>
	div {
		--width: 200px;
		position: relative;
		width: var(--width);
		margin: 10px;
		border: 1px solid #e0e0e0;
		background: #f6f6f6;
		padding: 0;
	}
	input {
		width: 100%;
		height: 100%;
		background: transparent;
		border: none;
		margin: 0;
		padding: 6px 8px;
		border-radius: 0;
	}
	ul {
		position: absolute;
		top: 100%;
		width: var(--width);
		left: -1px;
		display: none;
		padding: 14px 0;
		margin: 0;
		border: 1px solid #e0e0e0;
		background: #f6f6f6;
	}
	.expand ul {
		display: block;
	}
	li {
		display: none;
		padding: 8px;
		cursor: pointer;
	}
	.visible {
		display: block;
	}
	.indicator {
		background: #f0f0f0;
	}
</style>
<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import type { selectableData } from '../types';

	export let data: selectableData[] = [],
		placeholder = "",
		val = "",
		caseSensitive = false,
		locale = "en";

	const dispatch = createEventDispatcher();

	let selection = 0;
	let visibles: number[] = [];
	data.forEach( (_datum, i) => visibles.push(i));
	let focused = false;
	let hover = false;
	

	function inputAction(e: { key: any; }) {
		switch (e.key) {
			case "ArrowUp":
				// move up
				selection = (selection == 0) ? visibles.length - 1 : selection - 1;
				break;
				case "ArrowDown":
					selection = (selection == visibles.length - 1) ? 0 : selection + 1;
					// move down
				break;
			case "Enter":
				// move down
				handleSelection(visibles[selection], "keyboard");
				break;
		
			default:
				// search
				let arr: number[] = [];
				data.forEach( (datum, i) => {
					if (
							(caseSensitive && datum.text.includes(val)) ||
							(!caseSensitive && datum.text.toLocaleLowerCase(locale)
							 															.includes(val.toLocaleLowerCase(locale)))) {
						arr.push(i);
					}
				}
				)
				visibles = arr;
				selection = 0;
				break;
		}
	}
	function handleSelection(i: number, by: string) {
		dispatch("select", {
			i: Number(i),
			data: data[i],
			by,
		});
	}

</script>

<div
	class:expand={focused || hover}
	class="searchable-text-input-container"
	on:mouseover="{() => hover=true}"
	on:mouseout="{() => hover=false}"
	on:focus="{() => hover=true}"
  on:blur={() => hover=false}>
	<input
		class="searchable-text-input-input"
		{placeholder}
		on:focus="{()=>focused = true}"
		on:blur="{()=>focused = false}"
		bind:value={val}
		on:keyup={inputAction}
		type="text">
	<ul
		class="searchable-text-input-ul">
		{#each data as {text, attrs}, i}
			<li
				class="searchable-text-input-li"
				on:click="{(e) => handleSelection(i, "mouse")}"
				on:mouseover="{() => selection=i}"
        on:focus="{() => selection=i}"
				class:indicator={i==visibles[selection]}
				class:visible={visibles.includes(i)}
				{...attrs}
				>{text}</li>
		{/each}
	</ul>
</div>