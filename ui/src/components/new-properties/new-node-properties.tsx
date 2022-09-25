import React, {ChangeEvent} from "react";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import TextField from "@mui/material/TextField";
import DialogActions from "@mui/material/DialogActions";
import Button from "@mui/material/Button";
import {setModifiedState} from "../../utils/service";
import {getParsedModifiedState} from "../diagram-maker/helper/helper";

interface NewNodePropertiesProps {
    isOpen: boolean,
    nodeId: string,
    onClose: () => void,
}

export const NewNodeProperties = (props: NewNodePropertiesProps) => {
    let parsedModifiedState = getParsedModifiedState();

    const [payload, setPayload] = React.useState({
        name: parsedModifiedState.nodes[props.nodeId]?.consumerData["name"] !== undefined ? parsedModifiedState.nodes[props.nodeId].consumerData["name"] : "",
        type: parsedModifiedState.nodes[props.nodeId]?.consumerData["type"] !== undefined ? parsedModifiedState.nodes[props.nodeId].consumerData["type"] : "",
        language: parsedModifiedState.nodes[props.nodeId]?.consumerData["language"] !== undefined ? parsedModifiedState.nodes[props.nodeId].consumerData["language"] : "",
        isServer: parsedModifiedState.nodes[props.nodeId]?.consumerData["isServer"] !== undefined ? parsedModifiedState.nodes[props.nodeId].consumerData["isServer"] : false,
        isClient: parsedModifiedState.nodes[props.nodeId]?.consumerData["isClient"] !== undefined ? parsedModifiedState.nodes[props.nodeId].consumerData["isClient"] : false,
        // api resources to be generated
        resources: [],
        url: parsedModifiedState?.nodes[props.nodeId]?.consumerData["url"] !== undefined ? parsedModifiedState.nodes[props.nodeId].consumerData["url"] : "",
    });

    // TODO this is a hack as there is no NODE_UPDATE action in diagram-maker. We may later update this impl when we fork diagram-maker repo.
    // update state with additional properties added from UI (Post node creation)
    const handleUpdate = (event: React.MouseEvent<HTMLElement>) => {
        event.preventDefault();
        let parsedModifiedState = getParsedModifiedState();
        // update modifiedState with current fields on dialog box
        if (!(props.nodeId in parsedModifiedState.nodes)) {
            parsedModifiedState.nodes[props.nodeId] = {
                consumerData: {
                    type: payload.type,
                    name: payload.name,
                    isServer: payload.isServer,
                    isClient: payload.isClient,
                    language: payload.language,
                    url: payload.url
                }
            }
        } else {
            parsedModifiedState.nodes[props.nodeId].consumerData = {
                type: payload.type,
                name: payload.name,
                isServer: payload.isServer,
                isClient: payload.isClient,
                language: payload.language,
                url: payload.url
            }
        }
        // update modifiedState in the localstorage
        setModifiedState(JSON.stringify(parsedModifiedState))
        setPayload({
            ...payload,
            type: ""
        })
        props.onClose()
    }

    const handleTypeChange = (event: ChangeEvent<HTMLSelectElement | HTMLInputElement | HTMLTextAreaElement>) => {
        setPayload({
            ...payload,
            type: event.target.value
        });
    };

    const handleNameChange = (event: ChangeEvent<HTMLSelectElement | HTMLInputElement | HTMLTextAreaElement>) => {
        setPayload({
            ...payload,
            name: event.target.value
        });
    };

    const handleLanguageChange = (event: ChangeEvent<HTMLSelectElement | HTMLInputElement | HTMLTextAreaElement>) => {
        setPayload({
            ...payload,
            language: event.target.value
        });
    };

    const handleIsClientChange = (event: ChangeEvent<HTMLSelectElement | HTMLInputElement | HTMLTextAreaElement>) => {
        setPayload({
            ...payload,
            isClient: event.target.value
        });
    };

    const handleIsServerChange = (event: ChangeEvent<HTMLSelectElement | HTMLInputElement | HTMLTextAreaElement>) => {
        setPayload({
            ...payload,
            isServer: event.target.value
        });
    };

    const handleUrlChange = (event: ChangeEvent<HTMLSelectElement | HTMLInputElement | HTMLTextAreaElement>) => {
        setPayload({
            ...payload,
            url: event.target.value
        });
    };

    return <React.Fragment>
        <Dialog open={props.isOpen} onClose={props.onClose}>
            <DialogTitle>Node Properties : {props.nodeId}</DialogTitle>
            <DialogContent>
                <TextField
                    autoFocus
                    margin="dense"
                    id="name"
                    label="Name of Component"
                    type="text"
                    value={payload.name}
                    onChange={handleNameChange}
                    fullWidth
                    variant="standard"
                />
                <TextField
                    autoFocus
                    margin="dense"
                    id="type"
                    label="Type of Component"
                    type="text"
                    value={payload.type}
                    onChange={handleTypeChange}
                    fullWidth
                    variant="standard"
                />
                <TextField
                    autoFocus
                    margin="dense"
                    id="language"
                    label="Language"
                    type="text"
                    value={payload.language}
                    onChange={handleLanguageChange}
                    fullWidth
                    variant="standard"
                />
                <TextField
                    autoFocus
                    margin="dense"
                    id="isServer"
                    label="Is Server ?"
                    type="text"
                    value={payload.isServer}
                    onChange={handleIsServerChange}
                    fullWidth
                    variant="standard"
                />
                <TextField
                    autoFocus
                    margin="dense"
                    id="isClient"
                    label="Is Client ?"
                    type="text"
                    value={payload.isClient}
                    onChange={handleIsClientChange}
                    fullWidth
                    variant="standard"
                />
                <TextField
                    autoFocus
                    margin="dense"
                    id="url"
                    label="Url"
                    type="text"
                    value={payload.url}
                    onChange={handleUrlChange}
                    fullWidth
                    variant="standard"
                />
            </DialogContent>
            <DialogActions>
                <Button onClick={props.onClose}>Cancel</Button>
                <Button onClick={handleUpdate}>Update</Button>
            </DialogActions>
        </Dialog>
    </React.Fragment>;
}