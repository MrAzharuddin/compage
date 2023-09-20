import React from 'react';
import './App.scss';
import {BrowserRouter, Link, Route, Routes} from "react-router-dom";
import {Home} from "../home/home";
import Header from "../navbar/header";
import {Account} from "../account/account";
import {Grid} from "@mui/material";
import LoadingOverlay from "react-loading-overlay-ts";
import {useAppSelector} from "../../redux/hooks";
import {
    selectCreateProjectStatus, selectDeleteProjectStatus,
    selectExistsProjectStatus,
    selectGetProjectStatus,
    selectListProjectsStatus,
    selectUpdateProjectStatus
} from "../../features/projects-operations/slice";
import {selectGenerateCodeStatus} from "../../features/code-operations/slice";
import {selectUploadYamlStatus} from "../../features/open-api-yaml-operations/slice";
import {selectGetCurrentContextStatus} from "../../features/k8s-operations/slice";
import {CreateGitPlatform} from "../../features/git-platforms-operations/create-git-platform";
import {
    selectCreateGitPlatformStatus,
    selectDeleteGitPlatformStatus,
    selectListGitPlatformsStatus,
    selectUpdateGitPlatformStatus
} from "../../features/git-platforms-operations/slice";
import {SwitchProject} from "../../features/projects-operations/switch-project";
import {Login} from "../auth/login";
import {Projects} from "../../features/projects-operations/projects";
import {GitPlatforms} from "../../features/git-platforms-operations/git-platforms";

export const App = () => {
    const createProjectStatus = useAppSelector(selectCreateProjectStatus);
    const listProjectsStatus = useAppSelector(selectListProjectsStatus);
    const getProjectStatus = useAppSelector(selectGetProjectStatus);
    const existsProjectStatus = useAppSelector(selectExistsProjectStatus);
    const updateProjectStatus = useAppSelector(selectUpdateProjectStatus);
    const deleteProjectStatus = useAppSelector(selectDeleteProjectStatus);
    const createGitPlatformStatus = useAppSelector(selectCreateGitPlatformStatus);
    const listGitPlatformsStatus = useAppSelector(selectListGitPlatformsStatus);
    const updateGitPlatformStatus = useAppSelector(selectUpdateGitPlatformStatus);
    const deleteGitPlatformStatus = useAppSelector(selectDeleteGitPlatformStatus);
    const generateCodeStatus = useAppSelector(selectGenerateCodeStatus);
    const uploadYamlStatus = useAppSelector(selectUploadYamlStatus);
    const getCurrentContextStatus = useAppSelector(selectGetCurrentContextStatus);

    const isActive = () => {
        return createProjectStatus === 'loading'
            || listProjectsStatus === 'loading'
            || getProjectStatus === 'loading'
            || existsProjectStatus === 'loading'
            || updateProjectStatus === 'loading'
            || deleteProjectStatus === 'loading'
            || createGitPlatformStatus === 'loading'
            || listGitPlatformsStatus === 'loading'
            || updateGitPlatformStatus === 'loading'
            || deleteGitPlatformStatus === 'loading'
            || uploadYamlStatus === 'loading'
            || generateCodeStatus === 'loading'
            || getCurrentContextStatus === 'loading';
    };

    return <LoadingOverlay
        className="loading"
        active={isActive()}
        spinner
        text="Loading...">
        <BrowserRouter>
            <Grid
                container
                spacing={0}
                direction="row"
                justifyContent="center"
                alignItems="center"
                style={{
                    // TODO removes the scrollbar in the main screen
                    overflow: "auto",
                    width: "100%",
                    // TODO added 100% to take the whole webpage
                    // height: window.innerHeight - 150,
                    height: "100%",
                    // backgroundColor: 'teal'
                }}
            >
                <Header/>
                <br/>
                <Routes>
                    <Route path="/login" element={<Login/>}/>
                    <Route path="/create-git-platform" element={<CreateGitPlatform/>}/>
                    <Route path="/switch-project" element={<SwitchProject/>}/>
                    <Route path="/projects" element={<Projects/>}/>
                    <Route path="/git-platforms" element={<GitPlatforms/>}/>
                    <Route path="/home" element={<Home/>}/>
                    <Route path="/account" element={<Account/>}/>
                    <Route path="/" element={<Home/>}/>
                    <Route path="*" element={<p>Path not resolved <Link to={"/home"}> go to home</Link></p>}/>
                </Routes>
                {/*<Footer/>*/}
            </Grid>
        </BrowserRouter>
    </LoadingOverlay>;
};