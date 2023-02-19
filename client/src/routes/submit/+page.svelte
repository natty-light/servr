<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { loggedIn, selectedSchools, token } from '../../stores';
	import SearchableTextInput from '../../components/SearchableTextInput.svelte';
	import type { removeDetail, schoolOption, selectDetail } from '../../types';
	import SelectedSchool from '../../components/SelectedSchool.svelte';

	onMount(() => {
		if (!$loggedIn) {
			goto('/login');
		} else {
			if (!$token) {
				goto('/');
			}
		}
	});

	let selectableDataOptions: schoolOption[] = [
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
	let selected: schoolOption[] = [];
	const handleSelect = (event: { detail: selectDetail; }) => {
		const school = event.detail.option;
		selected = [...selected, school]
		selectableDataOptions = selectableDataOptions.filter( (option) => option.text != event.detail.option.text);
	}

	const handleRemove = (event: {detail: removeDetail; }) => {
		const school = event.detail.option
		selectableDataOptions = [...selectableDataOptions, school];
		selected = selected.filter((selectedSchool) => school.text != selectedSchool.text);
	}

	const handleSubmit = (e: MouseEvent & { currentTarget: EventTarget & HTMLButtonElement;}) => {
		e.preventDefault()
		selectedSchools.set(selected);
		goto('/view')
	}
</script>

<div class="flex flex-row">
	<div class="flex flex-col content-between w-6/12 items-center">
		<h1 class="text-lg mb-8">
			Your selected schools
		</h1>
		{#if selected.length > 0 }			
		<div class="border-2 rounded-md pl-4 pr-4 content-around">
			{#each selected as school, i }
			<SelectedSchool
			school={school}
			on:remove={handleRemove}
			/>
			{/each}
		</div>
		{/if}
	</div>
	<div class="flex flex-col content-between items-center w-6/12">
		<h1 class="text-lg mb-8">Select your schools</h1>
		<SearchableTextInput
		placeholder="Search..."
		data={selectableDataOptions}
		on:select={handleSelect} />
		<div class="mt-40">
			<button on:click={(e) => handleSubmit(e)}>
				Submit Selection
			</button>
		</div>
	</div>
</div>
	