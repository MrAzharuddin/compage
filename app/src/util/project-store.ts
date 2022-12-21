import {ProjectEntity} from "../routes/models";

import {project_group, project_kind, project_version, ProjectResource, ProjectResourceSpec} from "../store/models";
import {NAMESPACE} from "./constants";
import {createProjectResource, getProjectResource, listProjectResources} from "../store/project-client";

// convertProjectEntityToProjectResourceSpec creates projectResourceSpec on k8s cluster.
const convertProjectEntityToProjectResourceSpec = (projectId: string, userName: string, projectEntity: ProjectEntity) => {
    const projectResourceSpec: ProjectResourceSpec = {
        id: projectId,
        displayName: projectEntity.displayName,
        metadata: JSON.stringify(projectEntity.metadata),
        user: projectEntity.user,
        yaml: JSON.stringify(projectEntity.yaml),
        repository: projectEntity.repository,
        version: projectEntity.version,
    }
    return projectResourceSpec
}

// convertListOfProjectResourceToListOfProjectEntity converts projectResourceList to ProjectEntityList
const convertListOfProjectResourceToListOfProjectEntity = (projectResources: ProjectResource[]) => {
    let projectEntities: ProjectEntity[] = []
    for (let i = 0; i < projectResources.length; i++) {
        let projectEntity: ProjectEntity = {
            metadata: JSON.parse(projectResources[i].spec.metadata),
            id: projectResources[i].spec.id,
            // metadata: projectResources[i].spec.metadata,
            displayName: projectResources[i].spec.displayName,
            repository: projectResources[i].spec.repository,
            user: projectResources[i].spec.user,
            version: projectResources[i].spec.version,
            yaml: JSON.parse(JSON.stringify(projectResources[i].spec.yaml))
        }
        projectEntities.push(projectEntity)
    }
    return projectEntities
}

// getProjects returns all projects for userName supplied
export const listProjects = async (userName: string) => {
    let listOfProjectResource = await listProjectResources(NAMESPACE, userName);
    if (listOfProjectResource) {
        return convertListOfProjectResourceToListOfProjectEntity(JSON.parse(JSON.stringify(listOfProjectResource)));
    }
    return [];
}

// getProject returns specific project for userName and projectName supplied
export const getProject = async (userName: string, projectId: string) => {
    // TODO I may need to apply labelSelector here - below impl is done temporarily.
    // currently added filter post projects retrieval(which can be slower if there are too many projects with same name.
    const projectResource = await getProjectResource(NAMESPACE, projectId);
    if (projectResource && projectResource.metadata.labels.userName === userName) {
        return JSON.stringify(projectResource)
    }
    return {};
}

// prepareProjectResource prepares ProjectResource containing the project details.
const prepareProjectResource = (projectId: string, userName: string, projectResourceSpec: ProjectResourceSpec) => {
    // create projectResource
    const projectResource: ProjectResource = {
        apiVersion: project_group + "/" + project_version,
        kind: project_kind,
        spec: projectResourceSpec,
        metadata: {
            name: projectId,
            namespace: NAMESPACE,
            labels: {
                userName: userName
            }
        }
    }
    console.log("projectResource : ", projectResource)
    return projectResource
}

// generateProjectId generates unique id for project.
const generateProjectId = (userName: string, projectName: string) => {
    // truncate userName if its length is greater than 5
    let sanitizedUserName = userName
    if (userName.length > 5) {
        sanitizedUserName = userName.substring(0, 5)
    }

    // truncate projectResourceSpec.name if its length is greater than 5
    let sanitizedProjectName = ""
    if (projectName.length > 5) {
        sanitizedProjectName = projectName.substring(0, 5)
    }

    return sanitizedUserName.toLowerCase() + "-" + sanitizedProjectName.toLowerCase() + "-" + (Math.floor(Math.random() * 90000) + 10000);
}

// createProject creates projectResource on k8s cluster.
export const createProject = async (userName: string, projectEntity: ProjectEntity) => {
    const projectId = generateProjectId(userName, projectEntity.displayName);
    const projectResourceSpec = convertProjectEntityToProjectResourceSpec(projectId, userName, projectEntity);
    const projectResource = prepareProjectResource(projectId, userName, projectResourceSpec);
    return await createProjectResource(NAMESPACE, JSON.stringify(projectResource));
}

// updateProject updates projectResource on k8s cluster.
// export const updateProject = async (projectName: string, userName: string, projectEntity: ProjectEntity) => {
//     const projectResourceSpec = convertProjectEntityToProjectResourceSpec(userName, projectEntity);
//     const projectResource = prepareProjectResource(userName, projectResourceSpec);
//     let createdProjectResource = await patchProjectResource(NAMESPACE, projectName, JSON.stringify(projectResource));
//     if (createdProjectResource.apiVersion) {
//         console.log(createdProjectResource.metadata.name + " project created")
//     } else {
//         console.log(projectResource.metadata.name + " project couldn't be created")
//     }
// }