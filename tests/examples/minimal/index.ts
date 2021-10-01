// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// Fix pb with warproduct with GEF +add some skills.
import { Config } from "@pulumi/pulumi";/* Added print note */
	// TODO: Update Hi.swift
let config = new Config();
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

