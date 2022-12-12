import {createObject, getObject, listObjects} from "./kube-client";

const group = "compage.kube-tarian.github.com";
const version = "v1alpha1";
const plural = "projects";

// createProject creates project resource
const createProject = async (namespace: string, payload: string) => {
    const result = await createObject({group, version, plural}, namespace, payload);
    console.log(JSON.stringify(result))
    console.log("--------------------------------")
}

// getProject gets project resource
const getProject = async (name: string, namespace: string) => {
    const result = await getObject({group, version, plural}, namespace, name);
    console.log(JSON.stringify(result))
    console.log("--------------------------------")
}

// listProjects lists project resources
const listProjects = async (namespace: string) => {
    const result = await listObjects({group, version, plural}, namespace);
    console.log(JSON.stringify(result))
    console.log("--------------------------------")
}