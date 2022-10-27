import {getProjectGrpcClient} from "../grpc/project";
import {Router} from "express";

const grpcRouter = Router();

const projectGrpcClient = getProjectGrpcClient()
// generateProject (grpc calls to compage-core)
grpcRouter.post("/generate_project", async (req, res) => {
    const {repoName, yaml, projectName, userName} = req.body;
    try {
        const payload = {
            "name": projectName,
            "user": userName,
            "yaml": yaml,
            "repository": repoName
        }
        projectGrpcClient.GenerateProject(payload, (err: any, response: { fileChunk: any; }) => {
            if (err) {
                return res.status(500).json(err);
            }
            return res.status(200).json(response.fileChunk);
        });
    } catch (err) {
        return res.status(500).json(err);
    }
});

export default grpcRouter;
