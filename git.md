# gitlab CI
1. 目的:项目代码部署的流程
2. 利用流水线的方式将复杂的 部署流程 分 工序的写在:gitlabci.yml

```yaml
after_script:
  - pwd
  - cd ../;chown gitlab-runner:gitlab-runner -R $CI_PROJECT_NAME

stages: #定义所有的工序 所有工序按顺序执行 每个job同一个工序能够并发执行
#  - prepare
  - build
  - test
  - deploy
  - rsync

#prepare-docker:
#  stage: prepare
#  script:
#    - med prepare -n prepare
#    - med push -n prepare -t prepare

build-code:
  stage: build
  script:
    - med build -n build
    - med build -n release

validate:
  stage: test
  script:
    - med test -n validate

test:
  stage: test
  script:
    - med test -n test
  artifacts:
    paths:
      - ./med_export/coverage.html

pages: #git page 生成静态访问页面
  stage: deploy
  only:
    - develop
  dependencies:
    - test
  script:
    - mv med_export public
    - mv public/coverage.html public/index.html
  artifacts:
    paths:
      - public
    expire_in: 10 days

deploy-test:
  stage: deploy
  only:
    - develop-test
  script:
    - med update -n smart-sale-bg,smart-sale-bg-grpc-test,smart-sale-bg-grpcweb-test,smart-sale-bg-grpcweb-dev

deploy-staging:
  stage: deploy
  only:
    - develop
  script:
    - med update -n smart-sale-bg-grpcweb,smart-sale-bg-stage,smart-sale-bg-grpc

rsync-release:
  stage: rsync
  only:
    - tags
  script:
   - rm -rf build/
   - MED_VERSION=2 med cp release:/med/ build
   - chown gitlab-runner:gitlab-runner -R build
   - cd build && rm -rf go-configs && git clone git@git.guazi-corp.com:znkf-private/go-configs.git
   - tar -zcvf ../$CI_PROJECT_NAME-$CI_COMMIT_TAG.tar.gz * --exclude=.git --exclude=.idea && cd ../
   - chmod 755 $CI_PROJECT_NAME-$CI_COMMIT_TAG.tar.gz
   - rsync -avzp --progress $CI_PROJECT_NAME-$CI_COMMIT_TAG.tar.gz earthworm.guazi-corp.com::znkf/$CI_PROJECT_NAME/
   
   #推送打包好的可执行文件到 earthworm
```
