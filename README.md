# vector
This repository only exists to figure out how to publish a multi-package Go module
on Github using the git command. I'll be experimenting freely with it before i publish 
more serious stuff!

- On the github site, create a new repository, do not initialize.
- On the development machine, cd to the directory with the new module source code.
- git init
- git add .
- git commit -m "initial commit"
- git branch -M main
- git tag v0.1.0
- git remote add origin git@github.com:notnot/vector.git
- git push origin main
- git push origin v0.1.0

After editing the source code, it is time to publish a new version:

- git commit -a -m "an upgrade"
- git tag v0.1.1
- git push origin main
- git push origin v0.1.1



Perhaps i have not found the easiest way yet. I found out that tags have to be pushed
explicitly, otherwise only the main branch will be updated. Cumbersome.

You have to make sure that you have a SSH key set up on your development machine and
copied it to your github account (see settings). Otherwise an attempt to push will 
result in a "Permission denied (publickey)" error.
