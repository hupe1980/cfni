def cfni(synth):
    envs = {{.Envs}}

    for assembly in synth.assemblies:
        for stack in assembly.stacks:
            updated = False
            for func in stack.lambda_functions:
                if not func.has_envs(envs):
                    func.set_envs(envs)
                    stack.set_resource(func.id, func.resource)
                    updated = True
            
            if updated:
                synth.add_update(stack.absolute_filename, stack.template)

